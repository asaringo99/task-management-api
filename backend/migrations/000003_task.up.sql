-- Create Table
CREATE TABLE IF NOT EXISTS `tasks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `userid` int(11) NOT NULL,
    `boardid` int(11) NOT NULL,
    `priority` int(11) NOT NULL,
    `contents` VARCHAR(191) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (boardid) REFERENCES boards(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;