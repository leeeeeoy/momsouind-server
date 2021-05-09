package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/models"
)

type RecordDao struct {
}

var PublicRecordDao *RecordDao

func (*RecordDao) SelectByUserID(id int) *[]models.RecordData {
	records := []models.RecordData{}
	err := db.DBClient.Select(
		&records,
		"select * from record_datas where user_id = ?",
		id,
	)
	if err != nil {
		log.Println("녹음기록 조회x", err)
	}
	return &records
}
