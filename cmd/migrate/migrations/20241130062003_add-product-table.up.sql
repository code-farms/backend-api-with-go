CREATE TABLE IF NOT EXISTS products (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,    
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,
    `quantity` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`)
);