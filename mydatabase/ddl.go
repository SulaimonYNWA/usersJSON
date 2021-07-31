package mydatabase


const CreateUsersTable = `create table if not exists users(
	id integer Primary Key auto_increment,
	name varchar(20) not null,
	surname varchar(20) not null,
	age integer not null,
	sex varchar(20) not null,
	login varchar(20) unique,
	password varchar(20) not null,
	isAdmin boolean not null default false,
	remove boolean not null default false
	)`