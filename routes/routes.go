package routes

import (
	"testapi/handler/get/home"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/ss", home.Home)

	//get
	// e.GET("/", home.Home)
}
