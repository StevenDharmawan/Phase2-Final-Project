CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    password VARCHAR(200) NOT NULL,
    age INTEGER NOT NULL,
    role VARCHAR(100) NOT NULL DEFAULT 'customer'
);

CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    balance NUMERIC(10,2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE top_up_histories (
    id SERIAL PRIMARY KEY,
    wallet_id INTEGER NOT NULL,
    amount_top_up NUMERIC(10,2) NOT NULL,
    top_up_date_time TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);

CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    room_number VARCHAR(10) NOT NULL UNIQUE,
    category VARCHAR(100) NOT NULL,
    availability BOOLEAN NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    room_id INTEGER NOT NULL,
    checkin_date DATE NOT NULL,
    checkout_date DATE NOT NULL,
    total_price NUMERIC(10,2) NOT NULL
);