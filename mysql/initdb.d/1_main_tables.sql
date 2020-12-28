CREATE TABLE aapaca.artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    country VARCHAR(256) DEFAULT '' NOT NULL,
    birthday DATE,
    status INT NOT NULL, # 0 -> individual, 1 -> group
    image_url VARCHAR(256) DEFAULT 'http://placeimg.com/200/200/any' NOT NULL,
    description TEXT,
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
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES aapaca.artists(id)
);

CREATE TABLE aapaca.external_services (
    # service_id: amazon -> 1, apple -> 2, spotify -> 3
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE aapaca.external_ids (
    record_id INT NOT NULL,
    record_type INT NOT NULL, # record_type: artist -> 1, album -> 2, song -> 3
    external_id VARCHAR(256) NOT NULL,
    service_id INT NOT NULL,
    FOREIGN KEY (service_id) REFERENCES aapaca.external_services(id)
);