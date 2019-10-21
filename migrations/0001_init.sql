
create table users(
	id int AUTO_INCREMENT PRIMARY KEY,
	name varchar(100),
	username varchar(100),
	password varchar(100),
	email varchar(100),
	phone varchar(15),
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)


insert into users (name, username, email, phone, password) values ('albai','albai', 'albai@gmail.com','08989898', 'asdsad')