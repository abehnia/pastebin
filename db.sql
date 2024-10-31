DROP TABLE IF EXISTS pastebin;

CREATE TABLE IF NOT EXISTS pastebin (
    id uuid NOT NULL,
    "timestamp" bigint NOT NULL,
    title text NOT NULL,
    content TEXT NOT NULL,
    seen_counter INTEGER NOT NULL DEFAULT 0,
    star_counter INTEGER NOT NULL DEFAULT 0
);
