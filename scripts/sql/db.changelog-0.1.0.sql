-- liquibase formatted sql
-- changeset author:bpechiney

CREATE DATABASE IF NOT EXISTS challenge;
SET DATABASE = challenge;

-- TODO: remove this if you can find a workaround.
CREATE USER IF NOT EXISTS maxroach;
GRANT ALL ON DATABASE challenge TO maxroach;

DROP TABLE IF EXISTS challenge_user CASCADE;
CREATE TABLE challenge_user (
  id UUID PRIMARY KEY DEFAULT Gen_random_uuid(),
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  username VARCHAR(50) UNIQUE NOT NULL,
  role VARCHAR(21) CHECK (
    role='customer'
    OR role='administrator'
    OR role= 'attestation authority'
  ) NOT NULL,
  last_login TIMESTAMPTZ NOT NULL DEFAULT Statement_timestamp()
);
-- rollback DROP TABLE challenge_user CASCADE;

DROP TABLE IF EXISTS attestation CASCADE;
CREATE TABLE attestation (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	claimant_id UUID NOT NULL,
	claim VARCHAR(100) NOT NULL,
	INDEX (claimant_id),
	CONSTRAINT fk_attestation_challenge_user
		FOREIGN KEY (claimant_id)
		REFERENCES challenge_user (id)
);
-- rollback DROP TABLE attestation CASCADE;

DROP TABLE IF EXISTS challenge_user_attestation;
CREATE TABLE challenge_user_attestation (
  challenge_user_id UUID NOT NULL,
  attestation_id UUID NOT NULL,
  PRIMARY KEY (challenge_user_id, attestation_id),
  CONSTRAINT fk_challenge_user_attestation_challenge_user
    FOREIGN KEY (challenge_user_id)
    REFERENCES challenge_user (id),
  CONSTRAINT fk_challenge_user_attestation_attestation
    FOREIGN KEY (attestation_id)
    REFERENCES attestation (id)
);
-- rollback DROP TABLE challenge_user_attestation CASCADE;