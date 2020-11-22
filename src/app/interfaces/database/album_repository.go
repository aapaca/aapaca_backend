package database

import (
	"database/sql"
	"domain"
	"errors"
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

func (repo *AlbumRepository) GetAlbum(id int) (album domain.Album, err error) {
	// load album info
	pArtist := domain.Artist{}
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, artists.image_url, oc.id, oc.title,
								albums.id, albums.name, albums.released_date, albums.image_url, albums.description,
								albums.amazon_music_url, albums.apple_music_url, albums.spotify_url,
								p_art.id, p_art.name, p_art.image_url
							FROM albums
							INNER JOIN artists as p_art
								ON albums.primary_artist_id = p_art.id
								AND albums.id = ?
							LEFT OUTER JOIN contains
								ON albums.id = contains.album_id
							LEFT OUTER JOIN songs
								ON contains.song_id = songs.id
							LEFT OUTER JOIN performs
								ON songs.id = performs.song_id
							LEFT OUTER JOIN artists
								ON performs.artist_id = artists.id
							LEFT OUTER JOIN occupations as oc
								ON performs.occupation_id = oc.id
							`, id)
	defer rows.Close()
	creditMap := map[int]*domain.Credit{}
	var amazon, apple, spotify string
	var description sql.NullString
	var releasedDate sql.NullTime
	for rows.Next() {
		var nullableArtistID, partID sql.NullInt64
		var artistName, artistImgURL, partTitle sql.NullString
		if err = rows.Scan(&nullableArtistID, &artistName, &artistImgURL, &partID, &partTitle, &album.ID, &album.Name, &releasedDate, &album.ImageURL, &description, &amazon, &apple, &spotify, &pArtist.ID, &pArtist.Name, &pArtist.ImageURL); err != nil {
			return
		}
		if !nullableArtistID.Valid { // credit is empty
			break
		}
		artistID := int(nullableArtistID.Int64)
		if _, ok := creditMap[artistID]; !ok {
			creditMap[artistID] = &domain.Credit{
				Artist: domain.Artist{
					ID:       int(artistID),
					Name:     artistName.String,
					ImageURL: artistImgURL.String,
				},
				Parts: []domain.Occupation{},
			}
		}
		part := domain.Occupation{
			ID:    int(partID.Int64),
			Title: partTitle.String,
		}
		creditMap[artistID].Parts = append(creditMap[artistID].Parts, part)
	}
	// if rows have no columns, return err
	if album.ID != id {
		err = errors.New("album not found")
		return
	}
	album.PrimaryArtist = pArtist
	links := map[string]string{}
	if len(amazon) > 0 {
		links["amazonMusic"] = amazon
	}
	if len(apple) > 0 {
		links["appleMusic"] = apple
	}
	if len(spotify) > 0 {
		links["spotify"] = spotify
	}
	if len(links) > 0 {
		album.Links = links
	}
	if releasedDate.Valid {
		album.ReleasedDate = &releasedDate.Time
	}
	if description.Valid {
		album.Description = description.String
	}
	album.Credits = values(creditMap)
	return
}

func (repo *AlbumRepository) GetAlbumsByArtistId(artistId int) (albums []domain.Album, err error) {
	rows, err := repo.Query(`SELECT id, name, released_date, image_url FROM albums
							WHERE primary_artist_id = ?`, artistId)
	defer rows.Close()
	for rows.Next() {
		album := domain.Album{}
		var releasedDate sql.NullTime
		if err = rows.Scan(&album.ID, &album.Name, &releasedDate, &album.ImageURL); err != nil {
			return
		}
		if releasedDate.Valid {
			album.ReleasedDate = &releasedDate.Time
		}
		albums = append(albums, album)
	}
	if len(albums) == 0 {
		err = errors.New("albums not found")
		return
	}
	return
}
