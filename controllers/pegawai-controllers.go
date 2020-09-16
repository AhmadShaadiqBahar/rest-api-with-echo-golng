package controllers

import (
	"net/http"
	"strconv"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/models"
	"github.com/labstack/echo"
)

func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdatePegawai(c echo.Context) error {

	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		return c.JSON(http.StatusBadRequest, idErr)
	}
	// conv_id, err := strconv.Atoi(id)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	result, err := models.UpdatePegawai(userId, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		return c.JSON(http.StatusBadRequest, idErr)
	}

	result, err := models.DeletePegawai(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
