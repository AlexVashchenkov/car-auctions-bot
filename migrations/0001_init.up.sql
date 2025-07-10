-- USERS (участники)
CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    last_name   TEXT          NOT NULL,
    first_name  TEXT          NOT NULL,
    middle_name TEXT,
    phone       TEXT          NOT NULL UNIQUE,
    email       TEXT          NOT NULL,
    created_at  TIMESTAMP DEFAULT now()
);

-- ADMINS (админы, которые создают аукционы)
CREATE TABLE admins
(
    id          SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    last_name   TEXT          NOT NULL,
    first_name  TEXT          NOT NULL,
    middle_name TEXT,
    phone       TEXT          NOT NULL UNIQUE,
    email       TEXT          NOT NULL,
    created_at  TIMESTAMP DEFAULT now()
);

-- CARS (автомобили)
CREATE TABLE cars
(
    id            SERIAL PRIMARY KEY,
    make          TEXT NOT NULL,
    model         TEXT NOT NULL,
    year          INT  NOT NULL,
    mileage       INT,
    vin           TEXT,
    transmission  TEXT,
    drive_type    TEXT,
    fuel_type     TEXT,
    engine_volume NUMERIC,
    horsepower    INT,
    owners_count  INT,
    photo_urls    TEXT[], -- массив строк
    created_at    TIMESTAMP DEFAULT now()
);

-- AUCTIONS
CREATE TABLE auctions
(
    id               SERIAL PRIMARY KEY,
    car_id           INTEGER   NOT NULL REFERENCES cars (id) ON DELETE CASCADE,
    start_time       TIMESTAMP NOT NULL,
    duration_minutes INT       NOT NULL,
    starting_price   INT       NOT NULL,
    is_active        BOOLEAN   DEFAULT true,
    created_by       INTEGER REFERENCES admins (id),
    created_at       TIMESTAMP DEFAULT now()
);

-- BIDS (ставки)
CREATE TABLE bids
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    auction_id INTEGER REFERENCES auctions (id) ON DELETE CASCADE,
    amount     INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- WINNERS (победители прошедших аукционов)
CREATE TABLE winners
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id),
    auction_id INTEGER REFERENCES auctions (id) UNIQUE,
    final_bid  INT,
    won_at     TIMESTAMP DEFAULT now()
);
