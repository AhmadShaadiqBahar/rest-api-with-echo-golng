package routes

import (
	"net/http"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/controllers"
	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuth)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai/:user_id", controllers.UpdatePegawai)
	e.DELETE("/pegawai/:user_id", controllers.DeletePegawai)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	return e

}
