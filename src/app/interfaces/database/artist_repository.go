package database

import (
	"database/sql"
	"domain"
)

type ArtistRepository struct {
	SqlHandler
}

func generateArtistLinks(amazon string, apple string, spotify string) map[string]string {
	links := map[string]string{}
	if len(amazon) > 0 {
		links["amazonMusic"] = "https://www.amazon.com/" + amazon
	}
	if len(apple) > 0 {
		links["appleMusic"] = "https://music.apple.com/artist/" + apple
	}
	if len(spotify) > 0 {
		links["spotify"] = "https://open.spotify.com/artist/" + spotify
	}
	return links
}

func (repo *ArtistRepository) GetArtist(id int) (artist domain.Artist, err error) {
	rows, err := repo.Query(`SELECT artists.id, artists.name, artists.image_url, artists.description,
							artists.amazon_music_id, artists.apple_music_id, artists.spotify_id,
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
							WHERE artists.id = ?
							`, id, id, id)
	defer rows.Close()
	var members, aliases []domain.Artist
	var description sql.NullString
	var amazon, apple, spotify string
	for rows.Next() {
		var cartID sql.NullInt64
		var t, cartName, cartImgURL sql.NullString
		if err = rows.Scan(&artist.ID, &artist.Name, &artist.ImageURL, &description, &amazon, &apple, &spotify, &t, &cartID, &cartName, &cartImgURL); err != nil {
			return
		}
		if !cartID.Valid { // no aliases and members
			break
		}
		artist := domain.Artist{
			ID:       int(cartID.Int64),
			Name:     cartName.String,
			ImageURL: cartImgURL.String,
		}
		if t.String == "alias" {
			aliases = append(aliases, artist)
		} else { // type is "member"
			members = append(members, artist)
		}
	}
	if description.Valid {
		artist.Description = description.String
	}
	links := generateArtistLinks(amazon, apple, spotify)
	if len(links) > 0 {
		artist.Links = links
	}
	if len(aliases) > 0 {
		artist.Aliases = aliases
	}
	if len(members) > 0 {
		artist.Members = members
	}
	return
}
