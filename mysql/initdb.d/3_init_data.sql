INSERT INTO aapaca.external_services (name) VALUES ('amazon_music');
INSERT INTO aapaca.external_services (name) VALUES ('apple_music');
INSERT INTO aapaca.external_services (name) VALUES ('spotify');

INSERT INTO aapaca.artists (name, birthday, status, description) VALUES ('Common', '1972-03-13', 0, 'Lonnie Corant Jaman Shuka Rashid Lynn Jr. (born March 13, 1972), better known by his stage name Common (formerly Common Sense), is an American rapper, actor and writer.');
INSERT INTO aapaca.artists (name, birthday, status, description) VALUES ('Robert Glasper', '1978-04-05', 0, 'Robert Andre Glasper (born April 6, 1978) is an American pianist and record producer. He has been nominated for seven Grammy Awards and has won three Grammys and one Emmy Award.');
INSERT INTO aapaca.artists (name, birthday, status, description) VALUES ('Karriem Riggins', '1975-08-25', 0, 'Karriem Riggins is an American jazz drummer, hip hop producer, DJ and rapper.');
INSERT INTO aapaca.artists (name, birthday, status, description) VALUES ('August Greene', '2018-01-01', 1, 'August Greene is an American supergroup. The brainchild of rapper Common, and producers Robert Glasper and Karriem Riggins, it was formed in 2018 after sharing the Primetime Emmy Award for Outstanding Original Music and Lyrics for their song...');

INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 1, 'Common/e/B000AQ3K7O', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 1, '64490', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 1, '2GHclqNVjqGuiE5mA7BEoc', 3);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 1, 'Robert-Glasper/e/B00197K67U', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 1, '4643652', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 1, '5cM1PvItlR21WUyBnsdMcn', 3);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (3, 1, 'Karriem-Riggins/e/B0097A4F46', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (3, 1, '5421197', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (3, 1, '6e7BQ0gM6o8ecMXRZkXxlZ', 3);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (4, 1, 'August-Greene/e/B07BS2T1SC', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (4, 1, '1372391168', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (4, 1, '7KIFh8WDiqlz1Dlopd03cy', 3);

INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len) VALUES ('Joy and Peace', 1, 'Hip hop', '00:02:40');
INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len) VALUES ('Home', 1, 'Hip hop', '00:03:31');
INSERT INTO aapaca.songs (name, primary_artist_id, genre, song_len) VALUES ('Word From Moe Luv Interlude', 1, 'Hip hop', '00:00:40');

INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 3, 'B01MF58RNV', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 3, '1440895412?i=1440895418', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 3, '5VRq8bpbbSsROTLMTraLNN', 3);

INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date, description) VALUES ('Black America Again', 1, 'ARTium', '2016-11-04', 'Black America Again is the eleventh studio album by American rapper Common. It was released on November 4, 2016, by ARTium Recordings and Def Jam Recordings.');
INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date, description) VALUES ('August Greene', 4, 'August Greene LLC', '2018-05-04', '');
INSERT INTO aapaca.albums (name, primary_artist_id, label, released_date) VALUES ('Let Love', 1, 'Loma Vista', '2019-08-30');

INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 2, 'B01MCRSTN4', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 2, '1440895412', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (1, 2, '2DTbGcv40CSMW0P4yAQQBa', 3);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 2, 'B07B6HZ13D', 1);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 2, '1372391030', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (2, 2, '04o9bBHEVUGabrMRo3vzuy', 3);

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

INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES(1, 1, 1);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES(1, 2, 2);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES(1, 3, 3);

INSERT INTO aapaca.memberships (member_id, group_id) VALUES(1, 4);
INSERT INTO aapaca.memberships (member_id, group_id) VALUES(2, 4);
INSERT INTO aapaca.memberships (member_id, group_id) VALUES(3, 4);

INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(1, 1, 5);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(1, 2, 5);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(1, 3, 5);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(2, 3, 3);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(2, 3, 6);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(2, 3, 10);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(3, 1, 10);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(3, 2, 9);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(3, 2, 10);
INSERT INTO aapaca.performances(artist_id, song_id, occupation_id) VALUES(3, 3, 8);

INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(1, 1, 5);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(2, 1, 3);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(2, 1, 6);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(2, 1, 10);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(3, 1, 10);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(3, 1, 8);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(3, 1, 9);
INSERT INTO aapaca.participations(artist_id, album_id, occupation_id) VALUES(3, 1, 10);

INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Sheena Ringo', 'Japan', '1978-11-25', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Tokyo Incidents', 'Japan', '2004-09-08', 1);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Deyonna', 'Japan', '1978-11-25', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Rekishi', 'Japan', '1974-02-15',  0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Asai Kenichi', 'Japan', '1964-12-29', 0);
INSERT INTO aapaca.artists (name, country, birthday, status) VALUES ('Nagoshi Yukio', 'Japan', '1965-01-10', 0);

INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Honnou', 5);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tomei Ningen', 6);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Kira Kira Bushi', 8);
INSERT INTO aapaca.songs (name, primary_artist_id) VALUES ('Tsumi to Batsu', 5);

INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Shouso Strip', 5, '2000-03-31');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Adult', 6, '2006-01-25');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Rekitsu', 8, '2011-03-16');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Honnou', 5, '1999-10-27');
INSERT INTO aapaca.albums (name, primary_artist_id, released_date) VALUES ('Kouhukuron', 5, '1998-05-27');

INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (6, 2, '423534193', 2);
INSERT INTO aapaca.external_ids (record_id, record_type, external_id, service_id) VALUES (6, 2, '7l9bxNPPXK8458ScKSUXv6', 3);

INSERT INTO aapaca.occupations (title) VALUES ('Vocal');
INSERT INTO aapaca.occupations (title) VALUES ('Guitar');
INSERT INTO aapaca.occupations (title) VALUES ('Composer');
INSERT INTO aapaca.occupations (title) VALUES ('Lyricist');

INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES (4, 4, 12);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES (5, 5, 10);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES (6, 6, 2);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES (4, 7, 7);
INSERT INTO aapaca.contents (album_id, song_id, song_order) VALUES (7, 4, 1);

INSERT INTO aapaca.aliases (artist_id, alias_artist_id) VALUES (5, 7);
INSERT INTO aapaca.aliases (artist_id, alias_artist_id) VALUES (7, 5);

INSERT INTO aapaca.memberships (member_id, group_id) VALUES (5, 6);

INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 4, 1);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 5, 1);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (7, 6, 1);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (8, 6, 1);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 7, 1);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (9, 4, 2);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (10, 7, 2);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 4, 3);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (8, 6, 3);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 7, 3);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 4, 4);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 5, 4);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (8, 6, 4);
INSERT INTO aapaca.performances (artist_id, song_id, occupation_id) VALUES (5, 7, 4);

INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 4, 1);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 4, 3);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 4, 4);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 5, 1);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 5, 4);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 7, 1);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 7, 3);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (5, 7, 4);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (7, 6, 1);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (8, 6, 1);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (8, 6, 3);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (8, 6, 4);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (9, 4, 2);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (9, 7, 2);
INSERT INTO aapaca.participations (artist_id, album_id, occupation_id) VALUES (10, 4, 2);