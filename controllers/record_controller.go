package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/momsori-server/dao"
)

func GetRecord(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := dao.PublicRecordDao.SelectByUserID(id)
	return c.JSON(http.StatusOK, r)
}
