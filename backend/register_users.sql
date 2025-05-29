-- SQL для создания таблицы пользователей
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    sex ENUM('male','female') DEFAULT NULL,
    birth_date DATE DEFAULT NULL,
    nickname VARCHAR(50) UNIQUE,
    profile_image VARCHAR(255) DEFAULT NULL
);
