package repository

import (
	"database/sql"
	"domain"
	"interfaces/repository/rdb"
)

type AlbumRepository struct {
	rdb.SqlHandler
}

func (repo *AlbumRepository) GetAlbum(id int) (album domain.Album, err error) {
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, artists.image_url, oc.id, oc.title,
								albums.id, albums.name, albums.released_date, albums.image_url, albums.label, albums.description,
								external_ids.external_id, external_services.name,
								p_art.id, p_art.name, p_art.image_url
							FROM albums
							INNER JOIN artists as p_art
								ON albums.primary_artist_id = p_art.id
								AND albums.id = ?
							LEFT OUTER JOIN participations
								ON albums.id = participations.album_id
							LEFT OUTER JOIN artists
								ON participations.artist_id = artists.id
							LEFT OUTER JOIN occupations as oc
								ON participations.occupation_id = oc.id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = albums.id
								AND external_ids.record_type = ?
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							`, id, domain.RecordType.Album)
	if err != nil {
		return
	}
	defer rows.Close()

	pArtist := domain.Artist{}
	creditMap := map[int]*domain.Credit{}
	var description sql.NullString
	var releasedDate sql.NullTime
	links := domain.NewAlbumLinks()
	for rows.Next() {
		var nullableArtistID, partID sql.NullInt64
		var artistName, artistImgURL, partTitle, extID, extSName sql.NullString
		if err = rows.Scan(&nullableArtistID, &artistName, &artistImgURL, &partID, &partTitle, &album.ID, &album.Name, &releasedDate, &album.ImageURL, &album.Label, &description, &extID, &extSName, &pArtist.ID, &pArtist.Name, &pArtist.ImageURL); err != nil {
			return
		}
		if extID.Valid {
			err = links.AddLink(extID.String, extSName.String)
			if err != nil {
				return
			}
		}
		if !nullableArtistID.Valid { // credit is empty
			continue
		}
		artistID := int(nullableArtistID.Int64)
		if _, ok := creditMap[artistID]; !ok {
			creditMap[artistID] = &domain.Credit{
				Artist: &domain.Artist{
					ID:       int(artistID),
					Name:     artistName.String,
					ImageURL: artistImgURL.String,
				},
				Parts: &domain.Occupations{},
			}
		}
		if creditMap[artistID].Parts.Contains(int(partID.Int64)) {
			continue
		}
		part := domain.Occupation{
			ID:    int(partID.Int64),
			Title: partTitle.String,
		}
		creditMap[artistID].Parts.Append(part)
	}
	// if rows have no columns, album does not exist and album.Name becomes empty string.
	if album.Name == "" {
		return
	}
	album.PrimaryArtist = pArtist
	if links.Length() > 0 {
		album.Links = links
	}
	if releasedDate.Valid {
		album.ReleasedDate = &releasedDate.Time
	}
	if description.Valid {
		album.Description = description.String
	}
	for _, v := range creditMap {
		album.Credits = append(album.Credits, *v)
	}
	return
}

func (repo *AlbumRepository) GetAlbumsByArtistId(artistId int) (albums []domain.Album, err error) {
	rows, err := repo.Query(`SELECT id, name, released_date, image_url FROM albums
							WHERE primary_artist_id = ?`, artistId)
	if err != nil {
		return
	}
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
	return
}
