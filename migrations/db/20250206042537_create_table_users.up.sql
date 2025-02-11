CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `fullname` varchar(50) NOT NULL,
                         `email` varchar(255) NOT NULL,
                         `balance` double NOT NULL DEFAULT '0',
                         `api_username` VARCHAR(100) NOT NULL ,
                         `api_key` varchar(100) NOT NULL,
                         `suspend` boolean NOT NULL DEFAULT '0',
                         `is_admin` boolean NOT NULL DEFAULT '0',
                         `password_hash` varchar(255) NOT NULL,
                         `login_provider` enum('GOOGLE','REGISTER') NOT NULL DEFAULT 'REGISTER',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                     PRIMARY KEY(id),
                     UNIQUE KEY `email` (`email`),
                     UNIQUE KEY `api_username` (`api_username`),
                     UNIQUE KEY `api_key` (`api_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;