CREATE TABLE monobank_token (
    id int PRIMARY KEY,
    user_id in FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    token varchar(40) NOT NULL UNIQUE,
    expired_at varchar(40)
)