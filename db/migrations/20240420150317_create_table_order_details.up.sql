CREATE TABLE order_details(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    product_id INT NOT NULL ,
    order_id INT NOT NULL ,
    quantity INT NOT NULL ,
    total_price DECIMAL(10, 2) NOT NULL ,
    status INT NOT NULL COMMENT '1 = Pending, 2 = Accepted',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_products FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_orders FOREIGN KEY (order_id) REFERENCES orders(id),
    INDEX idx_created_at (created_at)
)ENGINE = InnoDB;