INSERT INTO artists (name, status, image_url) VALUES ('Artist 1', 0, 'http://www.example.com');
INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 1', 1, '00:02:40');
INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 2', 1, '00:00:40');
INSERT INTO songs (name, primary_artist_id, song_len) VALUES('Song 3', 1, '00:12:34');
INSERT INTO albums (name, primary_artist_id, released_date, image_url) VALUES('Album 1', 1, '2021-01-13', 'http://www.example.com');
INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 1, '1');
INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 2, '2');
INSERT INTO contents (album_id, song_id, song_order) VALUES (1, 3, '3');
INSERT INTO external_services (name) VALUE ('amazon_music');
INSERT INTO external_services (name) VALUE ('apple_music');
INSERT INTO external_services (name) VALUE ('spotify');
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', 'TEST1111', 1);
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', '1111', 2);
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (1, 'song', 'Test1111', 3);
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'song', 'TEST2222', 1);
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (2, 'song', '2222', 2);
INSERT INTO external_ids (record_id, record_type, external_id, service_id) VALUE (3, 'song', 'Test3333', 3);