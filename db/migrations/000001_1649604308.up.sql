CREATE TABLE IF NOT EXISTS users (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  first_name varchar(30) not null,
  last_name varchar(30) not null,
  username varchar(40) unique not null
);

CREATE TABLE IF NOT EXISTS monobank_token (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL,
    token varchar(40) NOT NULL UNIQUE,
    expired_at varchar(40),
    CONSTRAINT monobank_to_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);