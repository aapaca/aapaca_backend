CREATE TABLE songs (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    primary_artist_id INT NOT NULL,
    label_id INT,
    genre VARCHAR(256),
    PRIMARY KEY (id)
);