package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginUser struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

var (
	user = LoginUser{
		Name:     "Yoel",
		Email:    "hoheho18",
		Password: "1234",
	}
)

func Login(c echo.Context) error {
	var u LoginUser

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
	}

	if u.Email != user.Email || u.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	}

	token, err := CreateToken(u.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
