package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
)

type Userdatas struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Birth    string `json:"birth"`
}

var PublicUserDatas *Userdatas

func (*Userdatas) SelectTest(id int) *Userdatas {
	data := Userdatas{}
	err := db.DBClient.Get(&data,
		"select u.id, u.name, b.nickname, b.birth from users as u join babys as b where b.selected = 1 and u.id = ?", id)
	if err != nil {
		log.Println("조회 오류", err)
	}

	return &data
}
