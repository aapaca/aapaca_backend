INSERT INTO artists (name, status) VALUES ('Artist 1', 0);
INSERT INTO artists (name, status) VALUES ('Artist 2', 0);
INSERT INTO albums (name, primary_artist_id, label, released_date, image_url, description) VALUES('Album 1', 1, 'Label 1', '1999-07-13', 'http://www.example.com', 'This is test album 1');
INSERT INTO albums (name, primary_artist_id, label, released_date, image_url) VALUES('Album 2', 1, 'Label 2', '2021-01-13', 'http://www.example.com');