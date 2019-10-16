CREATE TABLE user (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    email VARCHAR(256) NOT NULL UNIQUE,
    name VARCHAR(32),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE question (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title TEXT,
    body TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- 初期データ
INSERT INTO user (email, name) VALUES ("hoge@hogehoge", "hoge");
INSERT INTO user (email, name) VALUES ("fuga@fugafuga", "foo");
INSERT INTO user (email, name) VALUES ("piyo@piyopiyo", "bar");

INSERT INTO question (title, body, user_id) VALUES ("How to use Git?", "What is 'git init?'", 1);
INSERT INTO question (title, body, user_id) VALUES ("How to use Docker?", "What is 'docker-compose?'", 2);