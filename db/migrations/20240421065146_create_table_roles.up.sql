CREATE TABLE roles(
    id INT AUTO_INCREMENT PRIMARY KEY ,
    roles VARCHAR(255) NOT NULL ,
    role_parent_id INT ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_roles (roles),
    INDEX idx_role_parent_id (role_parent_id),
    INDEX idx_created_at (created_at)
);