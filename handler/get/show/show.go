package show

import (
	"net/http"

	"testapi/db/conn"

	"github.com/labstack/echo/v4"
)

func Show(c echo.Context) error {
	d, err := conn.DbConnection()
	conn.ErrorCheck(err)
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"show": d,
	})
}
