CREATE DATABASE IF NOT EXISTS blog;

USE blog;

CREATE TABLE IF NOT EXISTS articles(
				id VARCHAR(50) NOT NULL,
				title VARCHAR(225) NOT NULL,
                author VARCHAR(225),
                content text,
				PRIMARY KEY (id)
				);

