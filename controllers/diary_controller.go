package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/momsori-server/dao"
)

func GetDiary(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := dao.PublicDiaryDao.SelectByUserID(id)
	return c.JSON(http.StatusOK, r)
}
