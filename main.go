package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leeeeeoy/momsori-server/db"
	"github.com/leeeeeoy/momsori-server/routes"
)

func main() {
	LoadEnv()
	db.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.UserRoutes(e)
	routes.RecordRoutes(e)
	routes.DiaryRoutes(e)

	fmt.Println("연결주소:", GetOutboundIP())
	e.Logger.Fatal(e.Start(":8080"))
}

//GetOutboundIP function
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.DB_HOST = os.Getenv("DB_HOST")
	db.DB_PORT = os.Getenv("DB_PORT")
	db.DB_USERNAME = os.Getenv("DB_USERNAME")
	db.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	db.DB_NAME = os.Getenv("DB_NAME")
}
