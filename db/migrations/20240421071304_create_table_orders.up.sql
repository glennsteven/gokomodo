CREATE TABLE orders(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    user_id INT NOT NULL ,
    delivery_source_address VARCHAR(255) NOT NULL ,
    delivery_destination_address VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_created_at (created_at),
    FOREIGN KEY (user_id) REFERENCES users(id)
);