CREATE TABLE artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    country VARCHAR(256) DEFAULT '' NOT NULL,
    birthday DATE,
    status INT NOT NULL, # 0 -> individual, 1 -> group
    image_url VARCHAR(256) DEFAULT 'http://placeimg.com/200/200/any' NOT NULL,
    description TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE occupations (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE albums (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256) DEFAULT '' NOT NULL,
    released_date DATE,
    image_url VARCHAR(256) DEFAULT 'http://placeimg.com/200/200/any' NOT NULL,
    description TEXT,
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES artists(id)
);

CREATE TABLE songs (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label VARCHAR(256) DEFAULT '' NOT NULL,
    genre VARCHAR(256) DEFAULT '' NOT NULL,
    song_len TIME,
    PRIMARY KEY (id),
    FOREIGN KEY (primary_artist_id) REFERENCES artists(id)
);

CREATE TABLE external_services (
    # service_id: amazon -> 1, apple -> 2, spotify -> 3
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE external_ids (
    record_id INT NOT NULL,
    record_type ENUM('artist', 'album', 'song') NOT NULL,
    external_id VARCHAR(256) NOT NULL,
    service_id INT NOT NULL,
    FOREIGN KEY (service_id) REFERENCES external_services(id)
);