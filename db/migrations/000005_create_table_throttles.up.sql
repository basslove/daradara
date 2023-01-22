CREATE TABLE IF NOT EXISTS throttles(
        hash_key TEXT NOT NULL,
        key_type TEXT NOT NULL,
        key TEXT NOT NULL,
        count INTEGER NOT NULL DEFAULT 0,
        count_expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
        block_expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY(hash_key, key_type)
);
