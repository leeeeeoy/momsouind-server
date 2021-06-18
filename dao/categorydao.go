package dao

import (
	"log"

	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/models"
)

type CategoryDao struct{}

var PublicCategoryDao *CategoryDao

func (*CategoryDao) InsertCategory(category models.Category) int64 {
	result := db.DBClient.MustExec(
		"insert into categories (name, user_id) values(?, ?)",
		category.Name, category.UserID)

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("카테고리 못넣음", err)
		return 0
	}
	return id
}
