CREATE TABLE IF NOT EXISTS sight_categories(
    id BIGSERIAL NOT NULL,
    name TEXT NOT NULL,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
COMMENT ON COLUMN "sight_categories"."name" IS 'free name';
COMMENT ON COLUMN "sight_categories"."is_valid" IS 'valid(true) or invalid(false)';
