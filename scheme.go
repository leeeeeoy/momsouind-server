package main

// var createUser = `create table users(
// 	id bigint not null auto_increment primary key,
// 	name varchar(10) not null,
// 	email varchar(45) not null,
// 	age varchar(10)  not null,
// 	password varchar(20) not null,
// 	baby_name varchar(10) not null,
// 	baby_nickname varchar(10) not null,
// 	baby_birth date not null,
// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 	)`

// var createRecordDatas = `
// 	create table record_datas(
// 		id bigint not null auto_increment primary key,
// 		FOREIGN KEY(user_id) REFERENCES uesrs (id),
// 		file_name varchar(20) not null,
// 		record_date date not null,
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)`

// var insertUser = `
// 	insert into users (name, email, age, password, baby_name, baby_nickname, baby_birth)
// 	values ("장요엘", "hoheho18@kakao.com","25","12345","엘슨", "엘슨엘슨", "1996-06-27")
// 	`

// var insertRecordData = `
// 	insert into record_datas (uesr_id, file_name, record_date)
// 	values (1, "test1", "2020-01-01")
// 	`
