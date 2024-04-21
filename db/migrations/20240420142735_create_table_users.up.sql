CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    role_enum INT NOT NULL COMMENT '1 = Pending, 2 = Accepted',
    email VARCHAR(100) NOT NULL ,
    name VARCHAR(255) NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    address VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_name (name),
    INDEX idx_email (email),
    INDEX idx_role_enum (role_enum),
    INDEX idx_created_at (created_at)
)ENGINE = InnoDB;