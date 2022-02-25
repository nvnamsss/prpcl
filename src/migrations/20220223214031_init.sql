-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE SCHEMA `new_schema` ;

CREATE TABLE `wagers` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `total_wager_value` FLOAT,
  `odds` INT NOT NULL,
  `selling_percentage` INT NOT NULL,
  `selling_price` DECIMAL(10,2) NOT NULL,
  `current_selling_price` DECIMAL(10,2) NOT NULL,
  `percentage_sold` FLOAT NOT NULL,
  `amount_sold` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `wager_transactions` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `wager_id` BIGINT(20) NOT NULL,
  `buying_price` DECIMAL(10,2) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
