CREATE TABLE
    IF NOT EXISTS accounts (
        id SERIAL PRIMARY KEY,
        user_id UUID REFERENCES auth.users,
        username TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );