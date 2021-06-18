package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
)

type FileDatasDao struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Files []file `json:"files"`
}
type file struct {
	Filename string `json:"filename"`
	Filedate string `json:"filedate"`
}
type datas struct {
	Name        string
	File_name   string
	Record_date string
}

var PublicFileDatasDao *FileDatasDao

func (*FileDatasDao) SelectAllByID(id int) *[]FileDatasDao {
	fileDatas := []FileDatasDao{}
	data := []datas{}
	err := db.DBClient.Select(&data,
		"select c.name, r.file_name, r.record_date from categories as c join record_datas as r on user_id = ? where c.id = r.category_id;", id)
	if err != nil {
		log.Println("조회 오류", err)
	}
	keys := make(map[string][]file)
	for i := 0; i < len(data); i++ {
		if _, ok := keys[data[i].Name]; ok {
			keys[data[i].Name] = append(keys[data[i].Name], file{
				Filename: data[i].File_name,
				Filedate: data[i].Record_date,
			})
		} else {
			keys[data[i].Name] = []file{
				{
					Filename: data[i].File_name,
					Filedate: data[i].Record_date,
				},
			}
		}
	}
	for key, val := range keys {
		fileDatas = append(fileDatas, FileDatasDao{
			Name:  key,
			Count: len(val),
			Files: val,
		})
	}
	return &fileDatas
}
