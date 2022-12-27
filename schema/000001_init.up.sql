CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    user_login VARCHAR(16) NOT NULL,
    user_hashed_password VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS api_info(
    user_id INT, 
    requests_count INT
);


INSERT INTO users (user_login, user_hashed_password) VALUES ('first_user', 'first_password');
INSERT INTO users (user_login, user_hashed_password) VALUES ('second_user', 'second_password');
INSERT INTO users (user_login, user_hashed_password) VALUES ('third_user', 'third_password');


INSERT INTO api_info (user_id, requests_count) VALUES(1, 10);