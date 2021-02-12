package database

import (
	"database/sql"
	"domain"
	"errors"
	"interfaces/database/rdb"
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
	rows, err := repo.Query(`SELECT artists.id, artists.name, artists.image_url, artists.description,
							external_ids.external_id, external_services.name,
							connections.type, c_art.id, c_art.name, c_art.image_url, oc.id, oc.title
							FROM artists
							LEFT OUTER JOIN (
								SELECT DISTINCT 'alias' AS type, alias_artist_id AS id, aliases.artist_id, credits.occupation_id
								FROM aliases
								LEFT OUTER JOIN (
									SELECT DISTINCT artist_id, occupation_id
									FROM performances
									UNION ALL
									SELECT DISTINCT artist_id, occupation_id
									FROM participations
								) AS credits
									ON alias_artist_id = credits.artist_id
								WHERE aliases.artist_id = ?
								UNION ALL
								SELECT DISTINCT 'member' AS type, member_id AS id, group_id AS artist_id, credits.occupation_id
								FROM memberships
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
								) AS credits
									ON member_id = credits.artist_id
							) AS connections
								ON artists.id = connections.artist_id
							LEFT OUTER JOIN occupations as oc
								ON connections.occupation_id = oc.id
							LEFT OUTER JOIN artists as c_art
								ON connections.id = c_art.id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = artists.id
								AND external_ids.record_type = ?
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							WHERE artists.id = ?
							`, id, id, id, domain.RecordType.Artist, id)
	defer rows.Close()
	memberMap := map[int]*domain.Credit{}
	aliasMap := map[int]*domain.Credit{}
	var description sql.NullString
	links := map[string]string{}
	for rows.Next() {
		var cartID, partID sql.NullInt64
		var t, cartName, cartImgURL, extID, extSName, partTitle sql.NullString
		if err = rows.Scan(&artist.ID, &artist.Name, &artist.ImageURL, &description, &extID, &extSName, &t, &cartID, &cartName, &cartImgURL, &partID, &partTitle); err != nil {
			return
		}
		if extID.Valid {
			c, l, e := generateArtistLink(extID.String, extSName.String)
			if err = e; err != nil {
				return
			}
			links[c] = l
		}
		if !cartID.Valid { // no aliases and members
			continue
		}
		aID := int(cartID.Int64)
		if t.String == "alias" {
			if _, ok := aliasMap[aID]; !ok {
				aliasMap[aID] = &domain.Credit{
					Artist: domain.Artist{
						ID:       aID,
						Name:     cartName.String,
						ImageURL: cartImgURL.String,
					},
					Parts: []domain.Occupation{},
				}
			}
			if !partID.Valid {
				continue
			}
			if partExists(aliasMap[aID].Parts, int(partID.Int64)) {
				continue
			}
			part := domain.Occupation{
				ID:    int(partID.Int64),
				Title: partTitle.String,
			}
			aliasMap[aID].Parts = append(aliasMap[aID].Parts, part)
		} else { // type is "member"
			if _, ok := memberMap[aID]; !ok {
				memberMap[aID] = &domain.Credit{
					Artist: domain.Artist{
						ID:       aID,
						Name:     cartName.String,
						ImageURL: cartImgURL.String,
					},
					Parts: []domain.Occupation{},
				}
			}
			if !partID.Valid {
				continue
			}
			if partExists(memberMap[aID].Parts, int(partID.Int64)) {
				continue
			}
			part := domain.Occupation{
				ID:    int(partID.Int64),
				Title: partTitle.String,
			}
			memberMap[aID].Parts = append(memberMap[aID].Parts, part)
		}
	}
	if description.Valid {
		artist.Description = description.String
	}
	if len(links) > 0 {
		artist.Links = links
	}
	if len(aliasMap) > 0 {
		aliases := []domain.Credit{}
		for _, v := range aliasMap {
			aliases = append(aliases, *v)
		}
		artist.Aliases = aliases
	}
	if len(memberMap) > 0 {
		members := []domain.Credit{}
		for _, v := range memberMap {
			members = append(members, *v)
		}
		artist.Members = members
	}
	return
}
