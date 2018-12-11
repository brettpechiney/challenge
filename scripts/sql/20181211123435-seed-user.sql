-- +migrate Up
INSERT INTO challenge.challenge_user (
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

-- +migrate Down
DELETE FROM challenge.challenge_user WHERE username = 'bpechiney';
