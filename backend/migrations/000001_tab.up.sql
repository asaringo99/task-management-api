-- Create Table
CREATE TABLE IF NOT EXISTS `tabs` (
    `id` int(11) NOT NULL,
    `userid` int(11) NOT NULL,
    `title` VARCHAR(191) NOT NULL,
    `is_active` BOOLEAN NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;