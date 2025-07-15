CREATE TABLE aggregated (
    id SERIAL PRIMARY KEY,                    -- Уникальный идентификатор (опционально, но удобно)
    pair_name TEXT NOT NULL,                 -- Название торговой пары, например "BTCUSDT"
    exchange TEXT NOT NULL,                  -- Название биржи, например "Binance"
    timestamp TIMESTAMP NOT NULL,            -- Время агрегации (обычно конец минуты)
    average_price DOUBLE PRECISION NOT NULL, -- Средняя цена
    min_price DOUBLE PRECISION NOT NULL,     -- Минимальная цена
    max_price DOUBLE PRECISION NOT NULL      -- Максимальная цена
);
