CREATE TABLE is_also_known_as (
    id INT NOT NULL AUTO_INCREMENT,
    artist_id INT NOT NULL,
    alias_artist_id INT NOT NULL,
    PRIMARY KEY (id)
);