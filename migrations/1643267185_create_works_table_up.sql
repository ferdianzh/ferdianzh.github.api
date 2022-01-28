CREATE TABLE IF NOT EXISTS works (
   "id" int NOT NULL AUTO_INCREMENT PRIMARY KEY,
   "id_role" int NOT NULL FOREIGN KEY REFERENCES "roles(id)",
   "name" varchar(50),
   "tech" varchar(50),
   "description" varchar(255),
) ENGINE=INNODB;

COMMIT;