-- +migrate Up
CREATE TABLE swift_codes (
    id TEXT PRIMARY KEY,
    country_iso2_code VARCHAR(2) NOT NULL,
    swift_code VARCHAR(11) NOT NULL UNIQUE,
    code_type VARCHAR(10) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    town_name VARCHAR(255),
    country_name VARCHAR(255) NOT NULL,
    time_zone VARCHAR(50)
);

CREATE INDEX idx_swift_left_8 ON swift_codes (LEFT(swift_code, 8));

-- +migrate Down
DROP INDEX IF EXISTS idx_swift_left_8;
DROP TABLE IF EXISTS swift_codes;