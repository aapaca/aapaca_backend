CREATE TABLE aapaca.contains (
    album_id INT NOT NULL,
    song_id INT NOT NULL,
    FOREIGN KEY (album_id) REFERENCES aapaca.albums(id),
    FOREIGN KEY (song_id) REFERENCES aapaca.songs(id)
);

CREATE TABLE aapaca.is_a_member_of (
    member_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (member_id) REFERENCES aapaca.artists(id),
    FOREIGN KEY (group_id) REFERENCES aapaca.artists(id)
);

CREATE TABLE aapaca.is_also_known_as (
    artist_id INT NOT NULL,
    alias_artist_id INT NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES aapaca.artists(id),
    FOREIGN KEY (alias_artist_id) REFERENCES aapaca.artists(id)
);

CREATE TABLE aapaca.performs (
    artist_id INT NOT NULL,
    song_id INT NOT NULL,
    occupation_id INT NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES aapaca.artists(id),
    FOREIGN KEY (song_id) REFERENCES aapaca.songs(id),
    FOREIGN KEY (occupation_id) REFERENCES aapaca.occupations(id)
);