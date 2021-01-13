CREATE TABLE contents (
    album_id INT NOT NULL,
    song_id INT NOT NULL,
    song_order VARCHAR(256) NOT NULL,
    FOREIGN KEY (album_id) REFERENCES albums(id),
    FOREIGN KEY (song_id) REFERENCES songs(id)
);

CREATE TABLE memberships (
    member_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (member_id) REFERENCES artists(id),
    FOREIGN KEY (group_id) REFERENCES artists(id)
);

CREATE TABLE aliases (
    artist_id INT NOT NULL,
    alias_artist_id INT NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    FOREIGN KEY (alias_artist_id) REFERENCES artists(id)
);

CREATE TABLE performances (
    artist_id INT NOT NULL,
    song_id INT NOT NULL,
    occupation_id INT NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    FOREIGN KEY (song_id) REFERENCES songs(id),
    FOREIGN KEY (occupation_id) REFERENCES occupations(id)
);

CREATE TABLE participations (
    artist_id INT NOT NULL,
    album_id INT NOT NULL,
    occupation_id INT NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    FOREIGN KEY (album_id) REFERENCES albums(id),
    FOREIGN KEY (occupation_id) REFERENCES occupations(id)
);