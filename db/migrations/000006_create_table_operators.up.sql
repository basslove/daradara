CREATE TYPE belong_types AS ENUM ('internal', 'external', 'unknown');
CREATE TYPE level_types AS ENUM ('normal', 'unknown');
CREATE TABLE IF NOT EXISTS operators(
    id BIGSERIAL NOT NULL,
    email TEXT NOT NULL UNIQUE,
    crypted_password TEXT NOT NULL,
    name TEXT NOT NULL,
    display_name TEXT NOT NULL DEFAULT '',
    image_url TEXT NOT NULL DEFAULT '',
    level level_types NOT NULL,
    belong belong_types NOT NULL,
    is_god BOOLEAN NOT NULL DEFAULT false,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    last_accessed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_logged_in_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

COMMENT ON COLUMN "operators"."name" IS 'full name';
COMMENT ON COLUMN "operators"."is_valid" IS 'valid(true) or invalid(false)';
