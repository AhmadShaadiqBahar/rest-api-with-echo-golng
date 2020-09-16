package routes

import (
	"net/http"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai/:user_id", controllers.UpdatePegawai)
	e.DELETE("/pegawai/:user_id", controllers.DeletePegawai)
	return e

}
