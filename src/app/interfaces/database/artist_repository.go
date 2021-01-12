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

func generateArtistLink(id string, serviceName string) (string, string, error) {
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
							connections.type, c_art.id, c_art.name, c_art.image_url
							FROM artists
							LEFT OUTER JOIN (
								SELECT 'alias' AS type, alias_artist_id AS id, artist_id
								FROM aliases
								WHERE artist_id = ?
								UNION ALL
								SELECT 'member' AS type, member_id AS id, group_id AS artist_id
								FROM memberships
								WHERE group_id = ?
							) AS connections
								ON artists.id = connections.artist_id
							LEFT OUTER JOIN artists as c_art
								ON connections.id = c_art.id
							LEFT OUTER JOIN external_ids
								ON external_ids.record_id = artists.id
								AND external_ids.record_type = ?
							LEFT OUTER JOIN external_services
								ON external_ids.service_id = external_services.id
							WHERE artists.id = ?
							`, id, id, domain.RecordType.Artist, id)
	defer rows.Close()
	memberMap := map[int]domain.Artist{}
	aliasMap := map[int]domain.Artist{}
	var description sql.NullString
	links := map[string]string{}
	for rows.Next() {
		var cartID sql.NullInt64
		var t, cartName, cartImgURL, extID, extSName sql.NullString
		if err = rows.Scan(&artist.ID, &artist.Name, &artist.ImageURL, &description, &extID, &extSName, &t, &cartID, &cartName, &cartImgURL); err != nil {
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
		cart := domain.Artist{
			ID:       int(cartID.Int64),
			Name:     cartName.String,
			ImageURL: cartImgURL.String,
		}
		if t.String == "alias" {
			aliasMap[cart.ID] = cart
		} else { // type is "member"
			memberMap[cart.ID] = cart
		}
	}
	if description.Valid {
		artist.Description = description.String
	}
	if len(links) > 0 {
		artist.Links = links
	}
	if len(aliasMap) > 0 {
		aliases := []domain.Artist{}
		for _, v := range aliasMap {
			aliases = append(aliases, v)
		}
		artist.Aliases = aliases
	}
	if len(memberMap) > 0 {
		members := []domain.Artist{}
		for _, v := range memberMap {
			members = append(members, v)
		}
		artist.Members = members
	}
	return
}
