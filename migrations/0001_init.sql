create table users(
	id int AUTO_INCREMENT PRIMARY KEY,
	username varchar(100),
	email varchar(100),
	phone varchar(15),
	createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
