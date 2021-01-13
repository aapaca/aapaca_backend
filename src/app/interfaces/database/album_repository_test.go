package database

import (
	"domain"
	"interfaces/database/rdb"
	"test/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func deleteAllRecords(sqlHandler rdb.SqlHandler) error {
	sqlHandler.Execute("SET FOREIGN_KEY_CHECKS = 0;")
	tables := []string{"aliases", "memberships", "contents", "performances", "participations", "external_services", "external_ids", "occupations", "artists", "songs", "albums"}
	for _, table := range tables {
		_, err := sqlHandler.Execute("TRUNCATE TABLE " + table + ";")
		if err != nil {
			return err
		}
	}
	sqlHandler.Execute("SET FOREIGN_KEY_CHECKS = 1;")
	return nil
}

func TestGetAlbum(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status, image_url) VALUES ('Test Artist 1', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Test Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Test Album 1', 1, 'Test Label', '2021-01-13', 'http://www.example.com', 'This is test album 1');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Test Album 2', 2, 'Test Label', '2021-01-13', 'http://www.example.com', 'This is test album 2');",
		"INSERT INTO occupations (title) VALUE ('Test Part 1')",
		"INSERT INTO occupations (title) VALUE ('Test Part 2')",
		"INSERT INTO external_services (name) VALUE ('amazon_music')",
		"INSERT INTO external_services (name) VALUE ('apple_music')",
		"INSERT INTO external_services (name) VALUE ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'album', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'album', 'Test2222', 3)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (1, 1, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (2, 1, 1)",
		"INSERT INTO participations (artist_id, album_id, occupation_id) VALUES (2, 1, 2)",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testArtist1 := domain.Artist{ID: 1, Name: "Test Artist 1", ImageURL: testURL}
	testArtist2 := domain.Artist{ID: 2, Name: "Test Artist 2", ImageURL: testURL}
	testPart1 := domain.Occupation{ID: 1, Title: "Test Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Test Part 2"}
	albumRepository := AlbumRepository{
		SqlHandler: sqlHandler,
	}
	t.Run("get all information of album", func(t *testing.T) {
		album, err := albumRepository.GetAlbum(1)
		expectedAlbum := domain.Album{
			ID:            1,
			Name:          "Test Album 1",
			PrimaryArtist: testArtist1,
			Credits: []domain.Credit{
				{Artist: testArtist1, Parts: []domain.Occupation{testPart1}},
				{Artist: testArtist2, Parts: []domain.Occupation{testPart1, testPart2}},
			},
			Label:        "Test Label",
			ReleasedDate: &testDate,
			ImageURL:     testURL,
			Description:  "This is test album 1",
			Links: map[string]string{
				"amazonMusic": "https://www.amazon.com/dp/TEST1111",
				"appleMusic":  "https://music.apple.com/album/1111",
				"spotify":     "https://open.spotify.com/album/Test1111",
			},
		}
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedAlbum, album, "Error")
	})
	t.Run("no credits", func(t *testing.T) {
		album, err := albumRepository.GetAlbum(2)
		expectedAlbum := domain.Album{
			ID:            2,
			Name:          "Test Album 2",
			PrimaryArtist: testArtist2,
			Label:         "Test Label",
			ReleasedDate:  &testDate,
			ImageURL:      testURL,
			Description:   "This is test album 2",
			Links: map[string]string{
				"spotify": "https://open.spotify.com/album/Test2222",
			},
		}
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedAlbum, album, "Error")
	})
	t.Run("empty album", func(t *testing.T) {
		emptyAlbum := domain.Album{}
		album, err := albumRepository.GetAlbum(100)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, emptyAlbum, album, "Error")
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAlbumByArtistId(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status) VALUES ('Test Artist 1', 0);",
		"INSERT INTO artists (name, status) VALUES ('Test Artist 2', 0);",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Test Album 1', 1, 'Test Label', '1999-07-13', 'http://www.example.com', 'This is test album 1');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url) VALUES('Test Album 2', 1, 'Test Label', '2021-01-13', 'http://www.example.com');",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	albumRepository := AlbumRepository{
		SqlHandler: sqlHandler,
	}
	testDate1, _ := time.Parse("2006-01-02", "1999-07-13")
	testDate2, _ := time.Parse("2006-01-02", "2021-01-13")
	testURL := "http://www.example.com"
	t.Run("Artist has 2 albums", func(t *testing.T) {
		testAlbum1 := domain.Album{ID: 1, Name: "Test Album 1", ReleasedDate: &testDate1, ImageURL: testURL}
		testAlbum2 := domain.Album{ID: 2, Name: "Test Album 2", ReleasedDate: &testDate2, ImageURL: testURL}
		expectedAlbums := []domain.Album{testAlbum1, testAlbum2}
		albums, err := albumRepository.GetAlbumsByArtistId(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedAlbums, albums, "Error")
	})
	t.Run("artist has no album", func(t *testing.T) {
		albums, err := albumRepository.GetAlbumsByArtistId(2)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 0, len(albums), "Error")
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}
