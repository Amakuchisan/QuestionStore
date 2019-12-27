CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    email VARCHAR(256) NOT NULL UNIQUE,
    name VARCHAR(32),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE questions (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    user_id INT NOT NULL REFERENCES user(id)
);

-- 初期データ
INSERT INTO users (email, name) VALUES ("hoge@hogehoge", "hoge");
INSERT INTO users (email, name) VALUES ("fuga@fugafuga", "foo");
INSERT INTO users (email, name) VALUES ("piyo@piyopiyo", "bar");

INSERT INTO questions (title, body, user_id) VALUES ("How to use Git?", "What is 'git init?'", 1);
INSERT INTO questions (title, body, user_id) VALUES ("How to use Docker?", "What is 'docker-compose?'", 2);