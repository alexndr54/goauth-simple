CREATE TABLE `reset_password` (
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    token varchar(255) NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;