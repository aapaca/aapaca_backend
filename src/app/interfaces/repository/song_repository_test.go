package repository

import (
	"domain"
	"interfaces/repository/rdb"
	"test"
	"test/infrastructure"
	"test/interfaces/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

type GetSongTestSuite struct {
	suite.Suite
	sqlHandler     rdb.SqlHandler
	songRepository SongRepository
}

func TestGetSongTestSuite(t *testing.T) {
	suite.Run(t, new(GetSongTestSuite))
}

func (suite *GetSongTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()
}

func (suite *GetSongTestSuite) SetupTest() {
	queries, err := repository.ReadSqlFile("testdata/get_song_init.sql")
	if err != nil {
		suite.T().Error(err)
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.songRepository = SongRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetSongTestSuite) TearDownTest() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetSongTestSuite) TestGetSong() {
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	testArtist1 := domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL}
	testArtist2 := domain.Artist{ID: 2, Name: "Artist 2", ImageURL: testURL}
	testPart1 := domain.Occupation{ID: 1, Title: "Part 1"}
	testPart2 := domain.Occupation{ID: 2, Title: "Part 2"}
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
	song, err := suite.songRepository.GetSong(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedSong.ID, song.ID)
	assert.Equal(suite.T(), expectedSong.Name, song.Name)
	assert.Equal(suite.T(), expectedSong.PrimaryArtist, song.PrimaryArtist)
	test.AssertCredits(suite.T(), expectedSong.Credits, song.Credits)
	assert.Equal(suite.T(), expectedSong.Album, song.Album)
	assert.Equal(suite.T(), expectedSong.SongLen, song.SongLen)
	assert.Equal(suite.T(), expectedSong.Genre, song.Genre)
	assert.Equal(suite.T(), expectedSong.Links, song.Links)
}

func (suite *GetSongTestSuite) TestGetSongNoCredit() {
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	testArtist1 := domain.Artist{ID: 1, Name: "Artist 1", ImageURL: testURL}
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
	song, err := suite.songRepository.GetSong(2)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), expectedSong, song)
}

func (suite *GetSongTestSuite) TestGetSongInvalidID() {
	emptySong := domain.Song{PrimaryArtist: domain.Artist{}, Album: domain.Album{}}
	song, err := suite.songRepository.GetSong(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), emptySong, song)
}

type GetAttendedSongsTestSuite struct {
	suite.Suite
	sqlHandler     rdb.SqlHandler
	songRepository SongRepository
}

func TestGetAttendedSongsTestSuite(t *testing.T) {
	suite.Run(t, new(GetAttendedSongsTestSuite))
}

func (suite *GetAttendedSongsTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()
}

func (suite *GetAttendedSongsTestSuite) SetupTest() {
	queries, err := repository.ReadSqlFile("testdata/get_attended_songs_init.sql")
	if err != nil {
		suite.T().Error(err)
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.songRepository = SongRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetAttendedSongsTestSuite) TearDownTest() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetAttendedSongsTestSuite) TestGetAttendedSongs() {
	// Song 1 and 2 are returned
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	expectedSongs := []domain.Song{
		domain.Song{ID: 1, Name: "Song 1", Album: testAlbum1},
		domain.Song{ID: 2, Name: "Song 2", Album: testAlbum1},
	}
	songs, err := suite.songRepository.GetAttendedSongs(2)
	if err != nil {
		suite.T().Error(err)
	}
	assert.ElementsMatch(suite.T(), expectedSongs, songs)
}

func (suite *GetAttendedSongsTestSuite) TestGetAttendedSongsAllSongs() {
	// all songs are returned
	testURL := "http://www.example.com"
	testDate, _ := time.Parse("2006-01-02", "2021-01-13")
	testAlbum1 := domain.Album{ID: 1, Name: "Album 1", ImageURL: testURL, ReleasedDate: &testDate}
	testAlbum2 := domain.Album{ID: 2, Name: "Album 2", ImageURL: testURL, ReleasedDate: &testDate}
	expectedSongs := []domain.Song{
		domain.Song{ID: 1, Name: "Song 1", Album: testAlbum1},
		domain.Song{ID: 2, Name: "Song 2", Album: testAlbum1},
		domain.Song{ID: 3, Name: "Song 3", Album: testAlbum2},
	}
	songs, err := suite.songRepository.GetAttendedSongs(3)
	if err != nil {
		suite.T().Error(err)
	}
	assert.ElementsMatch(suite.T(), expectedSongs, songs)
}

func (suite *GetAttendedSongsTestSuite) TestGetAttendedSongsNoSong() {
	// since Artist 1 is the primary artist for all song, songs is empty
	songs, err := suite.songRepository.GetAttendedSongs(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), 0, len(songs))
}

func (suite *GetAttendedSongsTestSuite) TestGetAttendedSongsInvalidID() {
	songs, err := suite.songRepository.GetAttendedSongs(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), 0, len(songs))
}

type GetSongsInAlbumTestSuite struct {
	suite.Suite
	sqlHandler     rdb.SqlHandler
	songRepository SongRepository
}

func TestGetSongsInAlbumTestSuite(t *testing.T) {
	suite.Run(t, new(GetSongsInAlbumTestSuite))
}

func (suite *GetSongsInAlbumTestSuite) SetupSuite() {
	suite.sqlHandler = infrastructure.NewSqlHandler()
}

func (suite *GetSongsInAlbumTestSuite) SetupTest() {
	queries, err := repository.ReadSqlFile("testdata/get_songs_in_album_init.sql")
	if err != nil {
		suite.T().Error(err)
	}
	for _, query := range queries {
		_, err := suite.sqlHandler.Execute(query)
		if err != nil {
			suite.T().Error(err)
		}
	}
	suite.songRepository = SongRepository{
		SqlHandler: suite.sqlHandler,
	}
}

func (suite *GetSongsInAlbumTestSuite) TearDownTest() {
	err := infrastructure.DeleteAllRecords(suite.sqlHandler)
	if err != nil {
		suite.T().Error(err)
	}
}

func (suite *GetSongsInAlbumTestSuite) TestGetSongsInAlbum() {
	// Album 1 contains 3 songs
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
	expectedSongs := []domain.Song{testSong1, testSong2, testSong3}
	songs, err := suite.songRepository.GetSongsInAlbum(1)
	if err != nil {
		suite.T().Error(err)
	}
	assert.ElementsMatch(suite.T(), expectedSongs, songs)
}

func (suite *GetSongsInAlbumTestSuite) TestGetSongsInAlbumInvalidID() {
	songs, err := suite.songRepository.GetSongsInAlbum(100)
	if err != nil {
		suite.T().Error(err)
	}
	assert.Equal(suite.T(), 0, len(songs))
}
