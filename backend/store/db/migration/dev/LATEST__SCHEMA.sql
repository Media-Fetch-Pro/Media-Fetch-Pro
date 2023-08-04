-- video_info
CREATE TABLE video_info (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    value TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    UNIQUE(name)
);


