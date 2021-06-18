package main

// var createUser = `create table users(
// 	id bigint not null auto_increment primary key,
// 	name varchar(10) not null,
// 	email varchar(45) not null,
// 	age varchar(10)  not null,
// 	password varchar(20) not null,
// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 	)`

// var createBaby = `create table babys(
// 		id bigint not null auto_increment primary key,
// 		user_id bigint not null,
// 		name varchar(10) not null,
// 		nickname varchar(10) not null,
// 		birth date not null,
// 		selected tinyint(1) not null,
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)`

// var createCategories = `
// 	create table categories(
// 		id bigint not null auto_increment primary key,
// 		name varchar(20) not null,
// 		user_id bigint not null,
// 		FOREIGN KEY(user_id) REFERENCES users (id),
// 		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)`

// var createRecordDatas = `
// create table record_datas(
// 	id bigint not null auto_increment primary key,
// 	category_id bigint not null,
// 	file_name varchar(20) not null,
// 	record_date date not null,
// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 	)`

// var createDiaries = `create table diaries(
// 	id bigint not null auto_increment primary key,
// 	user_id bigint not null,
// 	emotional_state int not null,
// 	physical_condition int not null,
// 	texts varchar(1000) not null,
// 	diary_date date not null,
// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 	)`

// var insertUser = `
// 	insert into users (name, email, age, password)
// 	values ("장요엘", "hoheho18@kakao.com","25","12345")
// 	`

// var insertBaby = `
// 	insert into babys (name, user_id, nickname, birth, selected)
// 	values ("장요엘", "hoheho18@kakao.com","25","12345","엘슨", "엘슨엘슨", "1996-06-27")
// 	`

// var insertCategory = `
// insert into categories (name, user_id)
// values ("1번 카테고리", 1)
// `

// var insertRecordData = `
// insert into record_datas (category_id, file_name, record_date)
// values (1, "test1", "2020-01-01")
// 	`

// var insertDiary = `
// 	insert into diary (user_id, emotional_state, physical_condition, texts, diary_date)
// 	values (1, 1, 1, "테스트 내용입니다.", "2020-01-01")
// `
