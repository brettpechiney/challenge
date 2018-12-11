-- +migrate Up
CREATE TABLE challenge.challenge_user (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  username VARCHAR(50) UNIQUE NOT NULL,
  password BYTES UNIQUE NOT NULL,
  role VARCHAR(21) CHECK (
    role='customer'
    OR role='administrator'
    OR role= 'attestation authority'
  ) NOT NULL,
  last_login TIMESTAMPTZ NOT NULL DEFAULT Statement_timestamp()
);

-- +migrate Down
DROP TABLE challenge.challenge_user;
