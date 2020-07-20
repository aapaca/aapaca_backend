CREATE TABLE aapaca.songs (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256),
    genre VARCHAR(256),
    PRIMARY KEY (id)
);

INSERT INTO songs (name, primary_artist_id) VALUES ("Automatic", 1)