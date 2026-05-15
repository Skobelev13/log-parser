CREATE TABLE IF NOT EXISTS logs (
    id UUID PRIMARY KEY,
    filename TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);