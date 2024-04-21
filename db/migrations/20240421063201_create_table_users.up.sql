CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    email VARCHAR(255) NOT NULL ,
    full_name VARCHAR(255) NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    address VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_email (email),
    INDEX idx_fullname (full_name),
    INDEX idx_password (password),
    INDEX idx_address (address),
    INDEX idx_created_at (created_at)
);