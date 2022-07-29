-- Создание таблицы
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL ,
    content TEXT NOT NULL,
    created DATETIME NOT NULL ,
    expires DATETIME NOT NULL
);

-- Добавление индекса для созданног столбца
CREATE INDEX idx_snippets_created ON snippets(created);

-- Добавлеение заметок
INSERT INTO snippets(title, content, created, expires)
    VALUES ('Подумав','Подумав — решайся,\nа решившись — не думай.', UTC_TIMESTAMP(),
            DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY )
            );

INSERT INTO snippets (title, content, created, expires)
VALUES ( 'Падая',
        'Семь раз упади, \nвосемь раз поднимись.',
        UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
        );


INSERT INTO snippets (title, content, created, expires)
VALUES ( 'Судьба',
        'Муж с женой должны быть подобны руке и глазам: \nкогда руке больно — глаза плачут, \nа когда глаза плачут — руки вытирают слезы.',
        UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
        );