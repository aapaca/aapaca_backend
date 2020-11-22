CREATE TABLE aapaca.artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    country VARCHAR(256) DEFAULT '' NOT NULL,
    birthday DATE,
    status INT NOT NULL, # 0 -> individual, 1 -> group
    image_url VARCHAR(256) DEFAULT 'http://placeimg.com/200/200/any' NOT NULL,
    description TEXT,
    amazon_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    apple_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    spotify_url VARCHAR(256) DEFAULT '' NOT NULL,
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
    label VARCHAR(256) DEFAULT '' NOT NULL,
    released_date DATE,
    image_url VARCHAR(256) DEFAULT 'http://placeimg.com/200/200/any' NOT NULL,
    description TEXT,
    amazon_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    apple_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    spotify_url VARCHAR(256) DEFAULT '' NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES aapaca.artists(id)
);

CREATE TABLE aapaca.songs (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256) DEFAULT '' NOT NULL,
    genre VARCHAR(256) DEFAULT '' NOT NULL,
    song_len TIME,
    amazon_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    apple_music_url VARCHAR(256) DEFAULT '' NOT NULL,
    spotify_url VARCHAR(256) DEFAULT '' NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES aapaca.artists(id)
);