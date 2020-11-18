INSERT INTO aapaca.artists (name, birthday, status, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('Common', '1972-03-13', 0, 'Lonnie Corant Jaman Shuka Rashid Lynn Jr. (born March 13, 1972), better known by his stage name Common (formerly Common Sense), is an American rapper, actor and writer.', 'https://www.amazon.com/Common/e/B000AQ3K7O/digital/ref=ntt_mp3_rdr?_encoding=UTF8&sn=d', 'https://music.apple.com/us/artist/common/64490', 'https://open.spotify.com/artist/2GHclqNVjqGuiE5mA7BEoc');
INSERT INTO aapaca.artists (name, birthday, status, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('Robert Glasper', '1978-04-05', 0, 'Robert Andre Glasper (born April 6, 1978) is an American pianist and record producer. He has been nominated for seven Grammy Awards and has won three Grammys and one Emmy Award.', 'https://www.amazon.com/Robert-Glasper/e/B00197K67U/digital/ref=ntt_mp3_rdr?_encoding=UTF8&sn=d', 'https://music.apple.com/us/artist/robert-glasper/4643652', 'https://open.spotify.com/artist/5cM1PvItlR21WUyBnsdMcn');
INSERT INTO aapaca.artists (name, birthday, status, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('Karriem Riggins', '1975-08-25', 0, 'Karriem Riggins is an American jazz drummer, hip hop producer, DJ and rapper.', 'https://www.amazon.com/Karriem-Riggins/e/B0097A4F46', 'https://music.apple.com/us/artist/karriem-riggins/5421197', 'https://open.spotify.com/artist/6e7BQ0gM6o8ecMXRZkXxlZ');
INSERT INTO aapaca.artists (name, birthday, status, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('August Greene', '2018-01-01', 1, 'August Greene is an American supergroup. The brainchild of rapper Common, and producers Robert Glasper and Karriem Riggins, it was formed in 2018 after sharing the Primetime Emmy Award for Outstanding Original Music and Lyrics for their song...', 'https://www.amazon.com/August-Greene/e/B07BS2T1SC/digital/ref=ntt_mp3_rdr?_encoding=UTF8&sn=d', 'https://music.apple.com/us/artist/august-greene/1372391168', 'https://open.spotify.com/artist/7KIFh8WDiqlz1Dlopd03cy');

INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len, amazon_music_url, apple_music_url, spotify_url) VALUES ('Joy and Peace', 1, 'Hip hop', '02:40', 'https://www.amazon.com/Joy-Peace-Explicit-feat-Bilal/dp/B01MF58RNV/ref=sr_1_2?dchild=1&keywords=common+joy+and+peace&qid=1603626572&sr=8-2', 'https://music.apple.com/album/joy-and-peace-feat-bilal/1440895412?i=1440895418', 'https://open.spotify.com/track/5VRq8bpbbSsROTLMTraLNN');
INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len) VALUES ('Home', 1, 'Hip hop', '03:31');
INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len) VALUES ('Word From Moe Luv Interlude', 1, 'Hip hop', '00:40');

INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('Black America Again', 1, 'ARTium', '2016-11-04', 'Black America Again is the eleventh studio album by American rapper Common. It was released on November 4, 2016, by ARTium Recordings and Def Jam Recordings.', 'https://www.amazon.com/Black-America-Again-Explicit-Common/dp/B01MCRSTN4', 'https://music.apple.com/us/album/black-america-again/1440895412', 'https://open.spotify.com/album/2DTbGcv40CSMW0P4yAQQBa');
INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date, description, amazon_music_url, apple_music_url, spotify_url) VALUES ('August Greene', 4, 'August Greene LLC', '2018-05-04', '', 'https://www.amazon.com/August-Greene-Amazon-Original/dp/B07B6HZ13D/ref=tmm_msc_swatch_0?_encoding=UTF8&qid=1605590686&sr=8-1', 'https://music.apple.com/album/august-greene/1372391030', 'https://open.spotify.com/album/04o9bBHEVUGabrMRo3vzuy');
INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date) VALUES ('Let Love', 1, 'Loma Vista', '2019-08-30');

INSERT INTO aapaca.occupations (title) VALUES ('Vocal');
INSERT INTO aapaca.occupations (title) VALUES ('Guitar');
INSERT INTO aapaca.occupations (title) VALUES ('Composer');
INSERT INTO aapaca.occupations (title) VALUES ('Lyricist');
INSERT INTO aapaca.occupations (title) VALUES ('Exective Producer');
INSERT INTO aapaca.occupations (title) VALUES ('Piano');
INSERT INTO aapaca.occupations (title) VALUES ('Keyboard');
INSERT INTO aapaca.occupations (title) VALUES ('Bass');
INSERT INTO aapaca.occupations (title) VALUES ('Drums');
INSERT INTO aapaca.occupations (title) VALUES ('Producer');

INSERT INTO aapaca.contains (album_id, song_id, song_order) VALUES(1, 1, 1);
INSERT INTO aapaca.contains (album_id, song_id, song_order) VALUES(1, 2, 2);
INSERT INTO aapaca.contains (album_id, song_id, song_order) VALUES(1, 3, 3);

INSERT INTO aapaca.is_a_member_of (member_id, group_id) VALUES(1, 4);
INSERT INTO aapaca.is_a_member_of (member_id, group_id) VALUES(2, 4);
INSERT INTO aapaca.is_a_member_of (member_id, group_id) VALUES(3, 4);

INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(1, 1, 5);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(1, 2, 5);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(1, 3, 5);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(2, 3, 3);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(2, 3, 6);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(2, 3, 10);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(3, 1, 10);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(3, 2, 9);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(3, 2, 10);
INSERT INTO aapaca.performs(artist_id, song_id, occupation_id) VALUES(3, 3, 8);

/*
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
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Kouhukuron', 1, '1998-05-27');

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
*/