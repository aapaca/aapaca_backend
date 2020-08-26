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
	pArtist := domain.Artist{
		Birthday: &defaultBirthday,
		Members:  defaultMembers,
		ImageURL: dummyURL,
	}

	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, oc.id, oc.title,
								albums.id, albums.name, albums.released_date, p_art.id, p_art.name
							FROM artists
							INNER JOIN performs
							 	ON performs.artist_id = artists.id
							INNER JOIN songs
								ON songs.id = performs.song_id
							INNER JOIN contains
								ON contains.song_id = songs.id
								AND contains.album_id = ?
							INNER JOIN occupations as oc
								ON oc.id = performs.occupation_id
							INNER JOIN albums
								ON albums.id = contains.album_id
							INNER JOIN artists as p_art
								ON p_art.id = albums.primary_artist_id`, id)
	defer rows.Close()
	creditMap := map[int]*domain.Credit{}
	for rows.Next() {
		var artistId int
		var artistName string
		part := domain.Occupation{}
		if err = rows.Scan(&artistId, &artistName, &part.ID, &part.Title, &album.ID, &album.Name, &album.ReleasedDate, &pArtist.ID, &pArtist.Name); err != nil {
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
	album.PrimaryArtist = pArtist
	album.ImageURL = dummyURL
	album.Credits = values(creditMap)
	return
}
