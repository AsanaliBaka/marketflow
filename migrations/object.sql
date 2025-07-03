CREATE TABLE crypto_prices (
    id            SERIAL PRIMARY KEY,         -- Уникальный ID (автоинкремент)
    pair_name     TEXT NOT NULL,              -- Название торговой пары
    exchange      TEXT NOT NULL,              -- Название биржи
    timestamp     TIMESTAMP NOT NULL,         -- Время сохранения
    average_price FLOAT8 NOT NULL,            -- Средняя цена за минуту
    min_price     FLOAT8 NOT NULL,            -- Минимальная цена за минуту
    max_price     FLOAT8 NOT NULL,            -- Максимальная цена за минуту

    UNIQUE (pair_name, exchange, timestamp)   -- Уникальность по этим трём полям
);
