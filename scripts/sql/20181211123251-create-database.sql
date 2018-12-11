-- +migrate Up notransaction
CREATE DATABASE challenge;

-- +migrate Down
DROP DATABASE challenge;