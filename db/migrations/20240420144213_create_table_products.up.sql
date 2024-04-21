CREATE TABLE products(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    product_name VARCHAR(255) NOT NULL ,
    description VARCHAR(255) NOT NULL ,
    price DECIMAL(10,2) NOT NULL ,
    user_id INT NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_product_name (product_name),
    INDEX idx_price (price),
    INDEX idx_created_at (created_at)
)ENGINE = InnoDB;