CREATE TABLE IF NOT EXISTS `accounts` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `owner` varchar(255) NOT NULL,
    `balance` bigint NOT NULL,
    `currency` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `entries` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `account_id` bigint NOT NULL,
    `amount` bigint NOT NULL COMMENT 'can be negative of positive',
    `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS `transfers` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `from_account_id` bigint NOT NULL,
    `to_account_id` bigint NOT NULL,
    `amount` bigint NOT NULL COMMENT 'must be positive',
    `created_at` timestamp NOT NULL DEFAULT (now())
);

-- ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

-- ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

-- ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_owner` ON `accounts` (`owner`);

CREATE INDEX `entries_index_account_id` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_from_account_id` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_to_account_id` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_from_and_to_account_id` ON `transfers` (`from_account_id`, `to_account_id`);