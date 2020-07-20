CREATE TABLE aapaca.albums (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256),
    released_date DATE,
    PRIMARY KEY (id)
);