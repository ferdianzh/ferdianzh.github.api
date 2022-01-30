CREATE TABLE IF NOT EXISTS works (
   id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
   id_role int NOT NULL,
   name varchar(50),
   tech varchar(50),
   description varchar(255),
   image varchar(50),
   FOREIGN KEY (id_role) REFERENCES roles(id)
) ENGINE=INNODB;

COMMIT;