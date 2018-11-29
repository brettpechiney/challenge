-- liquibase formatted sql
-- changeset author:bpechiney

CREATE DATABASE IF NOT EXISTS challenge;
SET DATABASE = challenge;

-- TODO: remove this if you can find a workaround.
CREATE USER IF NOT EXISTS maxroach;
GRANT ALL ON DATABASE challenge TO maxroach;

DROP TABLE IF EXISTS challenge_user CASCADE;
CREATE TABLE challenge_user (
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

INSERT INTO challenge_user (
  first_name,
  last_name,
  username,
  password,
  role
) VALUES (
  'Brett',
  'Pechiney',
  'bpechiney',
  '$2a$14$3yGTgnOiTW47DQo2PNHLi.KRCKpMfkg19trhrM5DjyWVwPgSAYFei',
  'administrator'
);
-- rollback DROP TABLE challenge_user CASCADE;

DROP TABLE IF EXISTS attestation CASCADE;
CREATE TABLE attestation (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	claimant_id UUID NOT NULL,
	attestor_id UUID NOT NULL,
	claim VARCHAR(100) NOT NULL,
	UNIQUE (claimant_id, attestor_id, claim),
	CONSTRAINT fk_attestation_challenge_user_claimant
		FOREIGN KEY (claimant_id)
		REFERENCES challenge_user (id),
	CONSTRAINT fk_attestation_challenge_user_attestor
		FOREIGN KEY (attestor_id)
		REFERENCES challenge_user (id)
);
-- rollback DROP TABLE attestation CASCADE;
