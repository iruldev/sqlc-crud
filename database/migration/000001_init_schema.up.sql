BEGIN;

CREATE TABLE accounts (
    id BIGINT NOT NULL AUTO_INCREMENT,
    owner VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL,
    currency VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now()),
    PRIMARY KEY (id)
);

CREATE TABLE entries (
   id BIGINT NOT NULL AUTO_INCREMENT,
   account_id BIGINT NOT NULL,
   amount BIGINT NOT NULL COMMENT 'can be negative or positive',
   created_at TIMESTAMP NOT NULL DEFAULT (now()),
   PRIMARY KEY (id)
--    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE transfers (
     id BIGINT NOT NULL AUTO_INCREMENT,
     from_account_id BIGINT NOT NULL,
     to_account_id BIGINT NOT NULL,
     amount BIGINT NOT NULL COMMENT 'must be positive',
     created_at TIMESTAMP NOT NULL DEFAULT (now()),
     PRIMARY KEY (id)
--      FOREIGN KEY (from_account_id) REFERENCES accounts(id),
--      FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

COMMIT;