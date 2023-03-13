CREATE TABLE IF NOT EXISTS `users` (
    `username` varchar(255) PRIMARY KEY,
    `hashed_password` varchar(255) NOT NULL,
    `full_name` varchar(255) NOT NULL,
    `email` varchar(255) UNIQUE NOT NULL,
    `password_changed_at` timestamp,
    `created_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

ALTER TABLE `accounts` ADD CONSTRAINT `owner_currency_key` UNIQUE (`owner`, `currency`);