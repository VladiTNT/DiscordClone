CREATE TABLE IF NOT EXISTS user (
    name TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,
    photo_id INTEGER,
    FOREIGN KEY(photo_id) REFERENCES photo(photo_id) ON DELETE CASCADE
);