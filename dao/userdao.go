package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/models"
)

type UserDao struct {
}

var PublicUserDao *UserDao

func (*UserDao) InsertUser(user *models.User) int64 {
	result := db.DBClient.MustExec(
		"insert into users (name, email, age, password) values(?, ?, ?, ?)",
		user.Name, user.Email, user.Age, user.Password)

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("유저 못넣음", err)
		return 0
	}
	return id
}

func (*UserDao) SelectByID(id int) *models.User {
	user := models.User{}
	err := db.DBClient.Get(&user, "select * from users where id = ?", id)
	if err != nil {
		log.Println("유저가 없음", err)
	}
	return &user
}

// func (*UserDao) UpdateByID(user *models.User) int64 {
// 	data, values, err := user.UpdateData()
// 	if err != nil {
// 		log.Println(err)
// 		return -1
// 	}

// 	result := db.DBClient.MustExec(
// 		fmt.Sprintf("update users set %s where id = ?", data),
// 		append(values, user.ID)...,
// 	)

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		log.Println(err)
// 		return -1
// 	}
// 	if rows != 1 {
// 		log.Println("잘못 변경됨; 변경된 갯수:", rows)
// 		return -1
// 	}

// 	return rows
// }

func (*UserDao) DeleteByID(id int) int64 {
	result := db.DBClient.MustExec(
		"delete from users where id = ?",
		id,
	)

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return -1
	}
	if rows != 1 {
		log.Println("잘못 삭제됨; 삭제된 갯수:", rows)
		return -1
	}
	return rows
}
