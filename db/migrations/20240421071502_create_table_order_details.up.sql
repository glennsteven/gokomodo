CREATE TABLE order_details(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    product_id INT NOT NULL ,
    order_id INT NOT NULL ,
    quantity INT NOT NULL ,
    total_price DECIMAL(10, 2) NOT NULL ,
    status INT DEFAULT 1 COMMENT '1: Pending, 2: Accepted',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_status (status),
    INDEX idx_created_at (created_at),
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);