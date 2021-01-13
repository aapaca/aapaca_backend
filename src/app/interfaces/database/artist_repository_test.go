package database

import (
	"domain"
	"sort"
	"test/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArtist(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Artist 1', 0, 'http://www.example.com', 'This is test artist 1')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Alias Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Group Artist 1', 1, 'http://www.example.com', 'This is test group artist 1')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Group Artist 2', 1, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url, description) VALUES ('Alias Group Artist 2', 1, 'http://www.example.com', 'This is test alias group artist 2')",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (2, 3);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (3, 2);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (5, 6);",
		"INSERT INTO aliases (artist_id, alias_artist_id) VALUES (6, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (1, 4);",
		"INSERT INTO memberships (member_id, group_id) VALUES (2, 4);",
		"INSERT INTO memberships (member_id, group_id) VALUES (1, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (2, 5);",
		"INSERT INTO memberships (member_id, group_id) VALUES (4, 5);",
		"INSERT INTO external_services (name) VALUE ('amazon_music')",
		"INSERT INTO external_services (name) VALUE ('apple_music')",
		"INSERT INTO external_services (name) VALUE ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'artist', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'artist', 'Test2222', 3)",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	artistRepository := ArtistRepository{
		SqlHandler: sqlHandler,
	}
	testURL := "http://www.example.com"
	t.Run("get all information of artist (not group, no aliases)", func(t *testing.T) {
		expectedArtist := domain.Artist{
			ID:          1,
			Name:        "Artist 1",
			ImageURL:    testURL,
			Description: "This is test artist 1",
			Links: map[string]string{
				"amazonMusic": "https://www.amazon.com/TEST1111",
				"appleMusic":  "https://music.apple.com/artist/1111",
				"spotify":     "https://open.spotify.com/artist/Test1111",
			},
		}
		artist, err := artistRepository.GetArtist(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedArtist, artist)
	})
	t.Run("not group, has aliases", func(t *testing.T) {
		expectedArtist := domain.Artist{
			ID:   2,
			Name: "Artist 2",
			Aliases: []domain.Artist{
				domain.Artist{ID: 3, Name: "Alias Artist 2", ImageURL: testURL},
			},
			ImageURL: testURL,
			Links: map[string]string{
				"spotify": "https://open.spotify.com/artist/Test2222",
			},
		}
		artist, err := artistRepository.GetArtist(2)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedArtist, artist)
	})
	t.Run("group, no aliases", func(t *testing.T) {
		expectedArtist := domain.Artist{
			ID:   4,
			Name: "Group Artist 1",
			Members: []domain.Artist{
				domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
			},
			Description: "This is test group artist 1",
			ImageURL:    testURL,
		}
		artist, err := artistRepository.GetArtist(4)
		if err != nil {
			t.Error(err)
		}
		// must fix
		assert.Equal(t, expectedArtist, artist)
	})
	t.Run("group, has aliases", func(t *testing.T) {
		expectedArtist := domain.Artist{
			ID:   5,
			Name: "Group Artist 2",
			Members: []domain.Artist{
				domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL},
				domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL},
				domain.Artist{ID: 4, Name: "Group Artist 1", ImageURL: testURL},
			},
			Aliases: []domain.Artist{
				domain.Artist{ID: 6, Name: "Alias Group Artist 2", ImageURL: testURL},
			},
			ImageURL: testURL,
		}
		artist, err := artistRepository.GetArtist(5)
		if err != nil {
			t.Error(err)
		}
		// Membersの順序がIDに対して昇順とは保障されていないのでソートする
		sort.Slice(artist.Members, func(i, j int) bool {
			return (*artist).Members[i].ID < artist.Members[j].ID
		})
		assert.Equal(t, expectedArtist, artist)
	})
	t.Run("Invalid ID", func(t *testing.T) {
		emptyArtist := domain.Artist{}
		artist, err := artistRepository.GetArtist(100)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, emptyArtist, artist)
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}
