package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJadwalscontrollers(c echo.Context) error {
	var jadwal []models.Jadwal
	if err := config.DB.Model(&jadwal).Debug().Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Find(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all jadwals", jadwal))
}
