/*
	Usage:
		First way:
			mysql -u root -p < <path_to_this_file>

		Second way:
			mysql -u root -p
			source <path_to_this_file>
*/ 

CREATE DATABASE IF NOT EXISTS url_shortener;
USE url_shortener;

CREATE TABLE IF NOT EXISTS user (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS url (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    url VARCHAR(2083) NOT NULL,
    click INT UNSIGNED DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
