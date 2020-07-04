CREATE TABLE albums (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL ,
    primary_artist VARCHAR(256) NOT NULL,
    attended_artist VARCHAR(256),
    label VARCHAR(256),
    released_date DATE,
    PRIMARY KEY (id)
);