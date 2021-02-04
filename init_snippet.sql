CREATE TABLE snippets (
id SERIAL NOT NULL PRIMARY KEY,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created TIMESTAMP NOT NULL,
expires TIMESTAMP NOT NULL
);

INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
TIMEZONE('utc', NOW()),
TIMEZONE('utc', NOW()) + INTERVAL '365 day');

INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
TIMEZONE('utc', NOW()),
TIMEZONE('utc', NOW()) + INTERVAL '365 day');

INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
TIMEZONE('utc', NOW()),
TIMEZONE('utc', NOW()) + INTERVAL '7 day');