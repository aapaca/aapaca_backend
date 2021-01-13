package database

import (
	"domain"
	"test/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShortenSongLen(t *testing.T) {
	t.Run("less than 10 minutes", func(t *testing.T) {
		in := "00:02:40"
		out := "2:40"
		got := shortenSongLen(in)
		assert.Equal(t, out, got)
	})
	t.Run("less than 1 minute", func(t *testing.T) {
		in := "00:00:15"
		out := "0:15"
		got := shortenSongLen(in)
		assert.Equal(t, out, got)
	})
	t.Run("more than 10 minutes", func(t *testing.T) {
		in := "00:12:34"
		out := "12:34"
		got := shortenSongLen(in)
		assert.Equal(t, out, got)
	})
	t.Run("more than 1 hour", func(t *testing.T) {
		in := "02:09:00"
		out := "2:09:00"
		got := shortenSongLen(in)
		assert.Equal(t, out, got)
	})
}

func TestGetSong(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 1', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO songs (name, primary_artist_id, genre, song_len) VALUES('Song 1', 1, 'Genre 1', '00:02:40');",
		"INSERT INTO songs (name, primary_artist_id, genre, song_len) VALUES('Song 2', 1, 'Genre 2', '00:00:40');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 1', 1, 'Label', '2021-01-13', 'http://www.example.com', 'This is test album 1');",
		"INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 2', 2, 'Label', '2021-01-13', 'http://www.example.com', 'This is test album 2');",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 1, '1')",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 2, '2')",
		"INSERT INTO occupations (title) VALUES ('Part 1')",
		"INSERT INTO occupations (title) VALUES ('Part 2')",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 1, 1)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 1, 2)",
		"INSERT INTO external_services (name) VALUES ('amazon_music')",
		"INSERT INTO external_services (name) VALUES ('apple_music')",
		"INSERT INTO external_services (name) VALUES ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUES (1, 'song', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUES (1, 'song', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUES (1, 'song', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUES (2, 'song', '2222', 2)",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	songRepository := SongRepository{
		SqlHandler: sqlHandler,
	}
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	testArtist1 := domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL}
	testArtist2 := domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL}
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
	t.Run("get all information of song", func(t *testing.T) {
		expectedSong := domain.Song{
			ID:            1,
			Name:          "Song 1",
			PrimaryArtist: testArtist1,
			Credits: []domain.Credit{
				{Artist: testArtist1, Parts: []domain.Occupation{testPart1, testPart2}},
				{Artist: testArtist2, Parts: []domain.Occupation{testPart2}},
			},
			Album:   testAlbum1,
			SongLen: "2:40",
			Genre:   "Genre 1",
			Links: map[string]string{
				"amazonMusic": "https://www.amazon.com/dp/TEST1111",
				"appleMusic":  "https://music.apple.com/album/1111",
				"spotify":     "https://open.spotify.com/track/Test1111",
			},
		}
		song, err := songRepository.GetSong(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedSong, song)
	})
	t.Run("no credits", func(t *testing.T) {
		expectedSong := domain.Song{
			ID:            2,
			Name:          "Song 2",
			PrimaryArtist: testArtist1,
			Album:         testAlbum1,
			SongLen:       "0:40",
			Genre:         "Genre 2",
			Links: map[string]string{
				"appleMusic": "https://music.apple.com/album/2222",
			},
		}
		song, err := songRepository.GetSong(2)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedSong, song)
	})
	t.Run("empty song", func(t *testing.T) {
		emptySong := domain.Song{PrimaryArtist: domain.Artist{}, Album: domain.Album{}}
		song, err := songRepository.GetSong(100)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, emptySong, song)
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAttendedSongs(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 1', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 2', 0, 'http://www.example.com')",
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 3', 0, 'http://www.example.com')",
		"INSERT INTO songs (name, primary_artist_id) VALUES('Song 1', 1);",
		"INSERT INTO songs (name, primary_artist_id) VALUES('Song 2', 1);",
		"INSERT INTO songs (name, primary_artist_id) VALUES('Song 3', 2);",
		"INSERT INTO albums (name, primary_artist_id, released_date, image_url) VALUES('Album 1', 1, '2021-01-13', 'http://www.example.com');",
		"INSERT INTO albums (name, primary_artist_id, released_date, image_url) VALUES('Album 2', 2, '2021-01-13', 'http://www.example.com');",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 1, '1')",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 2, '2')",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (2, 3, '1')",
		"INSERT INTO occupations (title) VALUES ('Part 1')",
		"INSERT INTO occupations (title) VALUES ('Part 2')",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 1, 1)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (1, 2, 1)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 2, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (2, 3, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (3, 1, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (3, 2, 2)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (3, 3, 1)",
		"INSERT INTO performances (artist_id, song_id, occupation_id) VALUES (3, 3, 2)",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	songRepository := SongRepository{
		SqlHandler: sqlHandler,
	}
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	testAlbum2 := domain.Album{ID: 2, Name: "Album 2", ImageURL: testURL, ReleasedDate: &testDate}
	t.Run("", func(t *testing.T) {
		expectedSongs := []domain.Song{
			domain.Song{ID: 1, Name: "Song 1", Album: testAlbum1},
			domain.Song{ID: 2, Name: "Song 2", Album: testAlbum1},
		}
		songs, err := songRepository.GetAttendedSongs(2)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, expectedSongs, songs)
	})
	t.Run("", func(t *testing.T) {
		expectedSongs := []domain.Song{
			domain.Song{ID: 1, Name: "Song 1", Album: testAlbum1},
			domain.Song{ID: 2, Name: "Song 2", Album: testAlbum1},
			domain.Song{ID: 3, Name: "Song 3", Album: testAlbum2},
		}
		songs, err := songRepository.GetAttendedSongs(3)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, expectedSongs, songs)
	})
	t.Run("", func(t *testing.T) {
		songs, err := songRepository.GetAttendedSongs(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 0, len(songs))
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSongsInAlbum(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	queries := []string{
		"INSERT INTO artists (name, status, image_url) VALUES ('Artist 1', 0, 'http://www.example.com')",
		"INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 1', 1, '00:02:40');",
		"INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 2', 1, '00:00:40');",
		"INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 3', 1, '00:12:34');",
		"INSERT INTO albums (name, primary_artist_id, released_date, image_url) VALUES('Album 1', 1, '2021-01-13', 'http://www.example.com');",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 1, '1')",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 2, '2')",
		"INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 3, '3')",
		"INSERT INTO external_services (name) VALUE ('amazon_music')",
		"INSERT INTO external_services (name) VALUE ('apple_music')",
		"INSERT INTO external_services (name) VALUE ('spotify')",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', 'TEST1111', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', '1111', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', 'Test1111', 3)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'song', 'TEST2222', 1)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'song', '2222', 2)",
		"INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (3, 'song', 'Test3333', 3)",
	}
	for _, query := range queries {
		_, err := sqlHandler.Execute(query)
		if err != nil {
			t.Error(err)
		}
	}
	songRepository := SongRepository{
		SqlHandler: sqlHandler,
	}
	testSong1 := domain.Song{
		ID:      1,
		Name:    "Song 1",
		SongLen: "2:40",
		Order:   "1",
		Links: map[string]string{
			"amazonMusic": "https://www.amazon.com/dp/TEST1111",
			"appleMusic":  "https://music.apple.com/album/1111",
			"spotify":     "https://open.spotify.com/track/Test1111",
		},
	}
	testSong2 := domain.Song{
		ID:      2,
		Name:    "Song 2",
		SongLen: "0:40",
		Order:   "2",
		Links: map[string]string{
			"amazonMusic": "https://www.amazon.com/dp/TEST2222",
			"appleMusic":  "https://music.apple.com/album/2222",
		},
	}
	testSong3 := domain.Song{
		ID:      3,
		Name:    "Song 3",
		SongLen: "12:34",
		Order:   "3",
		Links: map[string]string{
			"spotify": "https://open.spotify.com/track/Test3333",
		},
	}
	t.Run("", func(t *testing.T) {
		expectedSongs := []domain.Song{testSong1, testSong2, testSong3}
		songs, err := songRepository.GetSongsInAlbum(1)
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, expectedSongs, songs)
	})
	err := deleteAllRecords(sqlHandler)
	if err != nil {
		t.Error(err)
	}
}
