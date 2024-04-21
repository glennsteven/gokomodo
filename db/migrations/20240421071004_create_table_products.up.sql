CREATE TABLE products(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    user_id INT NOT NULL ,
    product_name VARCHAR(255) NOT NULL ,
    description TEXT ,
    price DECIMAL(10, 2) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_product_name (product_name),
    INDEX idx_price (price),
    INDEX idx_created_at (created_at),
    FOREIGN KEY (user_id) REFERENCES users(id)
);