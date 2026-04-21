CREATE TABLE `users` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	
	`email` VARCHAR(255) NOT NULL UNIQUE,
	`full_name` VARCHAR(255),
	`avatar` TEXT,
	`password` VARCHAR(255),
	`totp_secret` VARCHAR(255),
	`google_id`  VARCHAR(255),
	
	-- 	default luôn luôn có
	`deleted_at` TIMESTAMP NULL,
	`created_at` TIMESTAMP NOT NULL,
	`updated_at` TIMESTAMP NOT NULL
)

CREATE TABLE `foods` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	
	`name` VARCHAR(255) NOT NULL,
	`description`  VARCHAR(255),
	-- price TODO: migration
	
	-- 	default luôn luôn có
	`deleted_at` TIMESTAMP NULL,
	`created_at` TIMESTAMP NOT NULL,
	`updated_at` TIMESTAMP NOT NULL
)

CREATE TABLE `orders` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	
	`user_id` INT,
	`food_id` INT,
	 
	FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
	FOREIGN KEY (`food_id`) REFERENCES `foods`(`id`),
	
	-- 	default luôn luôn có
	`deleted_at` TIMESTAMP NULL,
	`created_at` TIMESTAMP NOT NULL,
	`updated_at` TIMESTAMP NOT NULL
)