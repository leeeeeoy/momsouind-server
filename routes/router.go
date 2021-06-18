package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/momsori-server/controllers"
)

func UserRoutes(e *echo.Echo) {
	// e.POST("/login", auth.Login)
	e.POST("/user", controllers.PostUser)
	e.GET("/user/:id", controllers.GetUserByID)
	// e.PATCH("/user", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	e.GET("/test/:id", controllers.SelectUserWithBabybyID)
}

func RecordRoutes(e *echo.Echo) {
	e.GET("/record/:id", controllers.GetRecordsByID)
}

func DiaryRoutes(e *echo.Echo) {
	e.GET("/diary/:id", controllers.GetDiaryByID)
}
func FileDatasRoute(e *echo.Echo) {
	e.GET("/files/:id", controllers.GetFiledatasByUserID)
}
