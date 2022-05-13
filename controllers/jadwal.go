package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetJadwalscontrollers(c echo.Context) error {
	var jadwal []models.Jadwal
	if err := config.DB.Model(&jadwal).Debug().Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Find(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all jadwal", jadwal))
}
func GetJadwalcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	jadwal := models.Jadwal{}
	if err := config.DB.First(&jadwal, id).Model(&jadwal).Debug().Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get jadwal", jadwal))
}
func GetHariJadwalHaricontrollers(c echo.Context) error {
	hari := c.Param("hari")
	jadwal := models.Jadwal{}
	if err := config.DB.Where("hari = ?", hari).Model(&jadwal).Debug().Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Find(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get jadwal by day", jadwal))
}

func CreateJadwalscontrollers(c echo.Context) error {
	jadwal := models.Jadwal{}
	c.Bind(&jadwal)

	if err := config.DB.Model(&jadwal).Debug().Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Debug().Create(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new dosen", jadwal))
}

func DeleteJadwalcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	jadwal := models.Jadwal{}
	if err := config.DB.Table("jadwal").First(&jadwal, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("jadwal deleted successfully", jadwal))
}

func UpdateJadwalcontrollers(c echo.Context) error {
	id := c.Param("id")
	jadwal := models.Jadwal{}

	if err := config.DB.First(&jadwal, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newjadwal := models.Jadwal{}
	c.Bind(&newjadwal)

	jadwal.Time = newjadwal.Time
	jadwal.Day = newjadwal.Day
	jadwal.IdMatakuliah = newjadwal.IdMatakuliah
	jadwal.NidDosen = newjadwal.NidDosen
	jadwal.IdRuangan = newjadwal.IdRuangan
	if err := config.DB.Save(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update matkul", jadwal))
}
