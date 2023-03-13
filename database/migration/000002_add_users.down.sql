ALTER TABLE `accounts` DROP CONSTRAINT `owner_currency_key`;

ALTER TABLE `accounts` DROP FOREIGN KEY `accounts_ibfk_1`;

DROP TABLE IF EXISTS users;