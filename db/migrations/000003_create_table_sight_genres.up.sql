CREATE TABLE IF NOT EXISTS sight_genres(
    id BIGSERIAL NOT NULL,
    sight_category_id BIGSERIAL NOT NULL,
    name TEXT NOT NULL,
    image_url TEXT NOT NULL DEFAULT '',
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(sight_category_id) REFERENCES sight_categories(id) ON DELETE CASCADE
);

COMMENT ON COLUMN "sight_genres"."name" IS 'free name';
COMMENT ON COLUMN "sight_genres"."is_valid" IS 'valid(true) or invalid(false)';
