CREATE TYPE gender_types AS ENUM ('male', 'female', 'unknown');
CREATE TYPE generation_types AS ENUM ('s10', 's20', 's30', 's40', 's50', 's60', 's70', 's80', 's90', 'unknown');
CREATE TABLE IF NOT EXISTS customers(
   id BIGSERIAL NOT NULL,
   email TEXT NOT NULL UNIQUE,
   crypted_password TEXT NOT NULL,
   name TEXT NOT NULL,
   gender gender_types NOT NULL,
   generation generation_types NOT NULL,
   display_name TEXT NOT NULL DEFAULT '',
   birthday TIMESTAMP WITH TIME ZONE NOT NULL,
   phone_number TEXT NOT NULL,
   introduction TEXT NOT NULL DEFAULT '',
   image_url TEXT NOT NULL DEFAULT '',
   allow_plan_displayed BOOLEAN NOT NULL DEFAULT false,
   is_valid BOOLEAN NOT NULL DEFAULT true,
   last_accessed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
   last_logged_in_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY(id)
);

COMMENT ON COLUMN "customers"."name" IS 'full name';
COMMENT ON COLUMN "customers"."is_valid" IS 'valid(true) or invalid(false)';
