INSERT INTO aapaca.artists (name, status) VALUES ('Sheena Ringo', 0);
INSERT INTO aapaca.artists (name, status) VALUES ('Tokyo Incidents', 0);
INSERT INTO aapaca.artists (name, status) VALUES ('Deyonna', 0);
INSERT INTO aapaca.artists (name, status) VALUES ('Rekishi', 0);

INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Honnou', 1);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tomei Ningen', 2);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Kira Kira Bushi', 4);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tsumi to Batsu', 1);

INSERT INTO aapaca.albums (name, primary_artist_id) VALUES ('Shouso Strip', 1);
INSERT INTO aapaca.albums (name, primary_artist_id) VALUES ('Adult', 2);
INSERT INTO aapaca.albums (name, primary_artist_id) VALUES ('Rekitsu', 4);
INSERT INTO aapaca.albums (name, primary_artist_id) VALUES ('Honnou', 1);

INSERT INTO aapaca.occupations (title) VALUES ('vocal');

INSERT INTO aapaca.contains (album_id, song_id) VALUES (1, 1);
INSERT INTO aapaca.contains (album_id, song_id) VALUES (2, 2);
INSERT INTO aapaca.contains (album_id, song_id) VALUES (3, 3);
INSERT INTO aapaca.contains (album_id, song_id) VALUES (1, 4);
INSERT INTO aapaca.contains (album_id, song_id) VALUES (4, 1);

INSERT INTO aapaca.is_also_known_as (artist_id, alias_artist_id) VALUES (1, 3);
INSERT INTO aapaca.is_also_known_as (artist_id, alias_artist_id) VALUES (3, 1);

INSERT INTO aapaca.is_a_member_of (member_id, group_id) VALUES (1, 2);

INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 1, 1);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 2, 1);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (3, 3, 1);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (4, 3, 1);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 4, 1);