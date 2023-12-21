-- Tabel Kompetisi
CREATE TABLE competition
(
    id   VARCHAR(64) PRIMARY KEY,
    date DATE
);

-- Tabel Skor
CREATE TABLE match
(
    id             SERIAL primary key,
    competition_id VARCHAR(64),
    team1_name     VARCHAR(64),
    team1_score    INT,
    team2_name     VARCHAR(64),
    team2_score    INT,
    FOREIGN KEY (competition_id) REFERENCES competition (id)
);
