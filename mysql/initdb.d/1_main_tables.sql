CREATE TABLE aapaca.artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    country VARCHAR(256),
    birthday DATE,
    status INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE aapaca.occupations (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE aapaca.albums (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256),
    released_date DATE,
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES aapaca.artists(id)
);

CREATE TABLE aapaca.songs (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256),
    genre VARCHAR(256),
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES aapaca.artists(id)
);