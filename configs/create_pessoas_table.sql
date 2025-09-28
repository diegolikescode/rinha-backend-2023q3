CREATE TABLE IF NOT EXISTS pessoas (
    id TEXT ,
    apelido TEXT PRIMARY KEY NOT NULL,
    nome TEXT,
    nascimento TEXT,
    stack TEXT,
    search_string TEXT,
    api_name TEXT
);

ALTER TABLE pessoas SET UNLOGGED;
