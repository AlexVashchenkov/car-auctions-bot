-- USERS
ALTER TABLE users
    ADD COLUMN state TEXT DEFAULT 'initial',
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

-- ADMINS
ALTER TABLE admins
    ADD COLUMN state TEXT DEFAULT 'initial',
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

-- CARS
ALTER TABLE cars
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

-- AUCTIONS
ALTER TABLE auctions
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

-- BIDS
ALTER TABLE bids
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();

-- WINNERS
ALTER TABLE winners
    ADD COLUMN updated_at TIMESTAMP DEFAULT now();
