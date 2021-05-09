package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/models"
)

type DiaryDao struct {
}

var PublicDiaryDao *DiaryDao

func (*DiaryDao) SelectByUserID(id int) *[]models.Diary {
	diarys := []models.Diary{}
	err := db.DBClient.Select(
		&diarys,
		"select * from diary where user_id = ?",
		id,
	)
	if err != nil {
		log.Println("다이어리 조회x", err)
	}
	return &diarys
}
