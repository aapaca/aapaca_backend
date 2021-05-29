package repository

import (
	"database/sql"
	"domain"
	"errors"
	"interfaces/repository/rdb"
)

type ArtistRepository struct {
	rdb.SqlHandler
}

func generateArtistLink(id, serviceName string) (string, string, error) {
	if serviceName == "amazon_music" {
		return "amazonMusic", "https://www.amazon.com/" + id, nil
	}
	if serviceName == "apple_music" {
		return "appleMusic", "https://music.apple.com/artist/" + id, nil
	}
	if serviceName == "spotify" {
		return "spotify", "https://open.spotify.com/artist/" + id, nil
	}
	return "", "", errors.New("invalid service name")
}

func (repo *ArtistRepository) GetArtist(id int) (artist domain.Artist, err error) {
	rows, err := repo.Query(`SELECT DISTINCT t.id, t.name, t.image_url, t.description, t.attr,
							t.occupation_id, oc.title, external_ids.external_id, external_services.name
							FROM
							(
								SELECT a.id, a.name, a.image_url, a.description, a.attr, p.occupation_id
								FROM
								(
									SELECT id, name, image_url, description, 'main' AS attr
									FROM artists
									WHERE id = ?
									UNION ALL
									SELECT id, name, image_url, '' AS description, 'alias' AS attr
									FROM aliases
									INNER JOIN artists
										ON artist_id = ?
										AND alias_artist_id = artists.id
								) a
								LEFT OUTER JOIN (
									SELECT performances.artist_id, performances.occupation_id
									FROM performances
									UNION ALL
									SELECT participations.artist_id, participations.occupation_id
									FROM participations
								) p
									ON p.artist_id = a.id
								UNION ALL
								SELECT id, name, image_url, '' AS description, 'member' AS attr, occupation_id
								FROM memberships
								INNER JOIN artists
									ON memberships.group_id = ?
									AND memberships.member_id = artists.id
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
									ON p.artist_id = id
							) t
							LEFT OUTER JOIN occupations as oc
								ON t.occupation_id = oc.id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = ?
								AND external_ids.record_type = ?
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							`, id, id, id, id, id, id, domain.RecordType.Artist)

	defer rows.Close()
	memberMap := map[int]*domain.Credit{}
	aliasMap := map[int]*domain.Credit{}
	links := map[string]string{}
	parts := []domain.Occupation{}
	for rows.Next() {
		var aID int
		var name, attr, imageURL string
		var ocID sql.NullInt64
		var desc, ocTitle, extID, extSName sql.NullString
		if err = rows.Scan(&aID, &name, &imageURL, &desc, &attr, &ocID, &ocTitle, &extID, &extSName); err != nil {
			return
		}
		if attr == "main" {
			artist.ID, artist.Name, artist.ImageURL = aID, name, imageURL
			if desc.Valid {
				artist.Description = desc.String
			}
			if extID.Valid {
				c, l, e := generateArtistLink(extID.String, extSName.String)
				if err = e; err != nil {
					return
				}
				links[c] = l
			}
			if !ocID.Valid {
				continue
			}
			if partExists(parts, int(ocID.Int64)) {
				continue
			}
			part := domain.Occupation{ID: int(ocID.Int64), Title: ocTitle.String}
			parts = append(parts, part)
		} else if attr == "member" {
			if _, ok := memberMap[aID]; !ok {
				memberMap[aID] = &domain.Credit{
					Artist: domain.Artist{ID: aID, Name: name, ImageURL: imageURL},
					Parts:  []domain.Occupation{},
				}
			}
			if !ocID.Valid {
				continue
			}
			if partExists(memberMap[aID].Parts, int(ocID.Int64)) {
				continue
			}
			part := domain.Occupation{ID: int(ocID.Int64), Title: ocTitle.String}
			memberMap[aID].Parts = append(memberMap[aID].Parts, part)
		} else { // attr == "alias"
			if _, ok := aliasMap[aID]; !ok {
				aliasMap[aID] = &domain.Credit{
					Artist: domain.Artist{ID: aID, Name: name, ImageURL: imageURL},
					Parts:  []domain.Occupation{},
				}
			}
			if !ocID.Valid {
				continue
			}
			if partExists(aliasMap[aID].Parts, int(ocID.Int64)) {
				continue
			}
			part := domain.Occupation{ID: int(ocID.Int64), Title: ocTitle.String}
			aliasMap[aID].Parts = append(aliasMap[aID].Parts, part)
		}
	}
	if len(links) > 0 {
		artist.Links = links
	}
	if len(parts) > 0 {
		artist.Parts = parts
	}
	if len(memberMap) > 0 {
		members := []domain.Credit{}
		for _, v := range memberMap {
			members = append(members, *v)
		}
		artist.Members = members
	}
	if len(aliasMap) > 0 {
		aliases := []domain.Credit{}
		for _, v := range aliasMap {
			aliases = append(aliases, *v)
		}
		artist.Aliases = aliases
	}
	return
}
