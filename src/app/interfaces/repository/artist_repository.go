package repository

import (
	"database/sql"
	"domain"
	"interfaces/repository/rdb"
)

type ArtistRepository struct {
	rdb.SqlHandler
}

func (repo *ArtistRepository) GetArtist(id int) (artist domain.Artist, err error) {
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, artists.image_url, artists.description,
							oc.id, oc.title, external_ids.external_id, external_services.name
							FROM artists
							LEFT OUTER JOIN (
								SELECT performances.artist_id, performances.occupation_id
								FROM performances
								UNION ALL
								SELECT participations.artist_id, participations.occupation_id
								FROM participations
							) p
								ON p.artist_id = artists.id
							LEFT OUTER JOIN occupations as oc
								ON p.occupation_id = oc.id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = artists.id
								AND external_ids.record_type = ?
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							WHERE artists.id = ?
							`, domain.RecordType.Artist, id)
	if err != nil {
		return
	}
	defer rows.Close()

	links := domain.NewArtistLinks()
	parts := domain.NewOccupations()
	for rows.Next() {
		var aID int
		var name, imageURL string
		var ocID sql.NullInt64
		var desc, ocTitle, extID, extSName sql.NullString
		if err = rows.Scan(&aID, &name, &imageURL, &desc, &ocID, &ocTitle, &extID, &extSName); err != nil {
			return
		}
		artist.ID, artist.Name, artist.ImageURL = aID, name, imageURL
		if desc.Valid {
			artist.Description = desc.String
		}
		if extID.Valid {
			err = links.AddLink(extID.String, extSName.String)
			if err != nil {
				return
			}
		}
		if !ocID.Valid || parts.Contains(int(ocID.Int64)) {
			continue
		}
		part := domain.Occupation{ID: int(ocID.Int64), Title: ocTitle.String}
		parts.Append(part)
	}
	if links.Length() > 0 {
		artist.Links = links
	}
	if !parts.IsEmpty() {
		artist.Parts = parts
	}
	return
}

func scanCredit(rows rdb.Row, creditMap map[int]*domain.Credit) {
	var aID int
	var name, imageURL string
	var ocID sql.NullInt64
	var ocTitle sql.NullString
	if err := rows.Scan(&aID, &name, &imageURL, &ocID, &ocTitle); err != nil {
		return
	}
	if _, ok := creditMap[aID]; !ok {
		creditMap[aID] = &domain.Credit{
			Artist: &domain.Artist{ID: aID, Name: name, ImageURL: imageURL},
			Parts:  domain.NewOccupations(),
		}
	}
	if !ocID.Valid {
		return
	}
	part := domain.Occupation{ID: int(ocID.Int64), Title: ocTitle.String}
	creditMap[aID].Parts.Append(part)
}

func (repo *ArtistRepository) FindMembers(id int) (members []domain.Credit, err error) {
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, artists.image_url, oc.id, oc.title
							FROM artists
							INNER JOIN memberships
								ON artists.id = member_id
								AND group_id = ?
							LEFT OUTER JOIN (
								SELECT performances.artist_id, performances.occupation_id
								FROM songs
								INNER JOIN performances
									ON songs.primary_artist_id = ?
									AND songs.id = performances.song_id
								UNION ALL
								SELECT participations.artist_id, participations.occupation_id
								FROM albums
								INNER JOIN participations
									ON albums.primary_artist_id = ?
									AND albums.id = participations.album_id
							) p
								ON p.artist_id = artists.id
							LEFT OUTER JOIN occupations as oc
								ON p.occupation_id = oc.id
							`, id, id, id)
	if err != nil {
		return
	}
	defer rows.Close()

	memberMap := map[int]*domain.Credit{}
	for rows.Next() {
		scanCredit(rows, memberMap)
	}
	for _, v := range memberMap {
		members = append(members, *v)
	}
	return
}

func (repo *ArtistRepository) FindAliases(id int) (aliases []domain.Credit, err error) {
	rows, err := repo.Query(`SELECT DISTINCT artists.id, artists.name, artists.image_url, oc.id, oc.title
							FROM artists
							INNER JOIN aliases
								ON artists.id = alias_artist_id
								AND artist_id = ?
							LEFT OUTER JOIN (
								SELECT performances.artist_id, performances.occupation_id
								FROM performances
								UNION ALL
								SELECT participations.artist_id, participations.occupation_id
								FROM participations
							) p
								ON p.artist_id = artists.id
							LEFT OUTER JOIN occupations as oc
								ON p.occupation_id = oc.id
							`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	aliasMap := map[int]*domain.Credit{}
	for rows.Next() {
		scanCredit(rows, aliasMap)
	}
	for _, v := range aliasMap {
		aliases = append(aliases, *v)
	}
	return
}
