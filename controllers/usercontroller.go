package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/momsori-server/dao"
	"github.com/leeeeeoy/momsori-server/models"
)

func PostUser(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	id := dao.PublicUserDao.InsertUser(&u)
	if id == 0 {
		log.Println("잘못넣었는데용")
	}

	return c.JSON(http.StatusCreated, id)
}

func GetUserByID(c echo.Context) error {
	u := &models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	u = dao.PublicUserDao.SelectByID(id)
	return c.JSON(http.StatusOK, &u)
}

func UpdateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	result := dao.PublicUserDao.UpdateByID(u)
	if result == -1 {
		log.Println("업데이트 안되는데용")
	}
	return c.JSON(http.StatusOK, "수정완료")
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := dao.PublicUserDao.DeleteByID(id)
	if result == -1 {
		log.Println("삭제 안됨요")
	}
	return c.JSON(http.StatusOK, "삭제완료")
}
