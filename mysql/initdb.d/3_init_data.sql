INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Sheena Ringo', 'Japan', '1978-11-25', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Tokyo Incidents', 'Japan', '2004-09-08', 1);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Deyonna', 'Japan', '1978-11-25', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Rekishi', 'Japan', '1974-02-15',  0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Asai Kenichi', 'Japan', '1964-12-29', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Nagoshi Yukio', 'Japan', '1965-01-10', 0);

INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Honnou', 1);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tomei Ningen', 2);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Kira Kira Bushi', 4);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tsumi to Batsu', 1);

INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Shouso Strip', 1, '2000-03-31');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Adult', 2, '2006-01-25');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Rekitsu', 4, '2011-03-16');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Honnou', 1, '1999-10-27');

INSERT INTO aapaca.occupations (title) VALUES ('Vocal');
INSERT INTO aapaca.occupations (title) VALUES ('Guitar');
INSERT INTO aapaca.occupations (title) VALUES ('Composer');
INSERT INTO aapaca.occupations (title) VALUES ('Lyricist');

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
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (5, 1, 2);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (6, 4, 2);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 1, 3);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (4, 3, 3);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 4, 3);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 1, 4);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 2, 4);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (4, 3, 4);
INSERT INTO aapaca.performs (artist_id, song_id, occupation_id) VALUES (1, 4, 4);