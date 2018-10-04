DROP TABLE IF EXISTS foo;
CREATE TABLE `foo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `foo` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`foo`)
);
