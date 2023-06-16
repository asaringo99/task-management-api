-- Create Table
CREATE TABLE IF NOT EXISTS `boards` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `userid` int(11) NOT NULL,
    `status` VARCHAR(191) NOT NULL,
    `priority` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;