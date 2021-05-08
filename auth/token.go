package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(useremail string) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "wow")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = useremail
	atClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
