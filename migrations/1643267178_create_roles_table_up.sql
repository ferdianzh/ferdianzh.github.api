CREATE TABLE IF NOT EXISTS roles (
   "id" int NOT NULL AUTO_INCREMENT PRIMARY KEY,
   "name" varchar(50),
   "description" varchar(255),
) ENGINE=INNODB;

COMMIT;