-- +goose Up
-- +goose StatementBegin
-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Asia/Jakarta";

-- Create books table
CREATE TABLE articles (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    title VARCHAR (255) NOT NULL,
    content TEXT NOT NULL,
    author UUID NOT NULL,
    likes INT NOT NULL DEFAULT 0,
    is_deleted BOOL NOT NULL DEFAULT false
);

-- Add indexes
CREATE INDEX active_articles ON articles (title) WHERE is_deleted = false;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
