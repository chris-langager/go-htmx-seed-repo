CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID UNIQUE,
    email VARCHAR(100) UNIQUE,
    hashed_password VARCHAR(100)
);