package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/models"
)

type BabyDao struct {
}

var PublicBabyDao *BabyDao

func (*BabyDao) InsertBaby(baby models.Baby) int64 {
	result := db.DBClient.MustExec(
		"insert into babys (user_id, name, nickname, birth, selected) values(?, ?, ?, ?, ?)",
		baby.UserID, baby.Name, baby.Nickname, baby.Birth, baby.Selected)

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("아기 못넣음", err)
		return 0
	}
	return id
}
func (*BabyDao) DeleteBabyByID(id int) int64 {
	result := db.DBClient.MustExec(
		"delete from babys where id = ?",
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
func (*BabyDao) SelectAll(userID int) *[]models.Baby {
	babys := []models.Baby{}
	err := db.DBClient.Select(&babys, "select * from babys where user_id = ?", userID)
	if err != nil {
		log.Println("아기가 없음", err)
	}
	return &babys
}
