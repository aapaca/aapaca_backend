CREATE TABLE artists (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    country VARCHAR(256),
    birthday DATE,
    PRIMARY KEY (id)
);