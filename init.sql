CREATE SCHEMA lastfm;

CREATE TABLE artists
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL UNIQUE,
	url VARCHAR ( 100 ) NOT NULL
);

CREATE TABLE albums
(
	id SERIAL PRIMARY KEY,
	title VARCHAR ( 50 ) NOT NULL,
    url VARCHAR ( 100 ) NOT NULL,
	artist_id INT NOT NULL,
    FOREIGN KEY (artist_id)
        REFERENCES artists (id)
);

CREATE TABLE tracks
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
	listeners BIGINT,
    playcount BIGINT,
    url VARCHAR ( 100 ) NOT NULL,
	album_id INT NOT NULL,
    FOREIGN KEY (album_id)
        REFERENCES albums (id)
);

CREATE TABLE tags
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL UNIQUE,
	url VARCHAR ( 100 ) NOT NULL
);

CREATE TABLE tracks_tag
(
	id SERIAL PRIMARY KEY,
	track_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (track_id)
        REFERENCES tracks (id),
    FOREIGN KEY (tag_id)
        REFERENCES tags (id),
    UNIQUE (track_id, tag_id)
);


