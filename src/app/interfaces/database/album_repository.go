package database

import (
	"domain"
	"time"
)

type AlbumRepository struct {
	SqlHandler
}

func values(creditMap map[int]*domain.Credit) []domain.Credit {
	vs := []domain.Credit{}
	for _, v := range creditMap {
		vs = append(vs, *v)
	}
	return vs
}

func (repo *AlbumRepository) FindById(id int) (album domain.Album, err error) {
	const dummyURL = "http://placeimg.com/200/200/any"
	defaultBirthday := time.Time{}
	defaultMembers := []domain.Artist{}
	// load album info
	row, err := repo.Query(`SELECT albums.id, albums.name, artists.id, artists.name, albums.released_date
							FROM albums
							INNER JOIN artists
								ON albums.primary_artist_id = artists.id
								AND albums.id = ?`, id)
	defer row.Close()
	if err != nil {
		return
	}
	pArtist := domain.Artist{
		Birthday: &defaultBirthday,
		Members:  defaultMembers,
		ImageURL: dummyURL,
	}
	row.Next()
	if err = row.Scan(&album.ID, &album.Name, &pArtist.ID, &pArtist.Name, &album.ReleasedDate); err != nil {
		return
	}
	album.PrimaryArtist = pArtist
	album.ImageURL = dummyURL

	// load credits
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, occupations.id, occupations.title
							FROM artists
							INNER JOIN performs
							 	ON performs.artist_id = artists.id
							INNER JOIN songs
								ON songs.id = performs.song_id
							INNER JOIN contains
								ON contains.song_id = songs.id
								AND contains.album_id = ?
							INNER JOIN occupations
							 	ON occupations.id = performs.occupation_id`, id)
	defer rows.Close()
	creditMap := map[int]*domain.Credit{}
	for rows.Next() {
		var artistId int
		var artistName string
		part := domain.Occupation{}
		if err = rows.Scan(&artistId, &artistName, &part.ID, &part.Title); err != nil {
			return
		}
		if _, ok := creditMap[artistId]; !ok {
			creditMap[artistId] = &domain.Credit{
				Artist: domain.Artist{
					ID:       artistId,
					Name:     artistName,
					Birthday: &defaultBirthday,
					Members:  defaultMembers,
					ImageURL: dummyURL,
				},
				Parts: []domain.Occupation{},
			}
		}
		creditMap[artistId].Parts = append(creditMap[artistId].Parts, part)
	}
	album.Credits = values(creditMap)
	return
}
