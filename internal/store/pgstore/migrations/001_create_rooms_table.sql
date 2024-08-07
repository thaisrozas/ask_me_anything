DROP TABLE IF EXISTS rooms;

-- Create the table
CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    theme VARCHAR(255) NOT NULL
);