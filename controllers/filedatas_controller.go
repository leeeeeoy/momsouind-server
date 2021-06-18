package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/momsori-server/dao"
)

func GetFiledatasByUserID(c echo.Context) error {
	filedatas := &[]dao.FileDatasDao{}
	id, _ := strconv.Atoi(c.Param("id"))
	filedatas = dao.PublicFileDatasDao.SelectAllByID(id)
	return c.JSON(http.StatusOK, filedatas)
}
