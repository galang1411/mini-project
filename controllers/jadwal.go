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

func GetHariJadwalHaricontrollers(c echo.Context) error {
	hari := c.Param("hari")
	jadwal := []models.Jadwal{}
	if err := config.DB.Model(&jadwal).Debug().Where("hari", hari).Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Find(&jadwal).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get jadwal", jadwal))
}

func GetJadwalIDcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	jadwal := models.Jadwal{}
	if err := config.DB.Model(&jadwal).Debug().Where("id", id).Preload("Matakuliah").Preload("Dosen").Preload("Ruangan").Find(&jadwal).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get jadwal", jadwal))
}

//
func CreateJadwalscontrollers(c echo.Context) error {
	jadwal := models.InsertJadwal{}
	c.Bind(&jadwal)

	if err := config.DB.Table("jadwal").Create(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new dosen", jadwal))
}

func DeleteJadwalcontrollers(c echo.Context) error {
	id := c.Param("id")
	jadwal := models.InsertJadwal{}

	if err := config.DB.Table("jadwal").Where("id", id).Find(&jadwal).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Table("jadwal").Where("id", id).Delete(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("jadwal deleted successfully", jadwal))
}

func UpdateJadwalcontrollers(c echo.Context) error {
	id := c.Param("id")
	jadwal := models.InsertJadwal{}

	if err := config.DB.Table("jadwal").Where("id", id).Find(&jadwal).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "jadwal not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newjadwal := models.InsertJadwal{}
	c.Bind(&newjadwal)

	jadwal.Time = newjadwal.Time
	jadwal.Hari = newjadwal.Hari
	jadwal.Dosen = newjadwal.Dosen
	jadwal.Ruangan = newjadwal.Ruangan
	jadwal.Matakuliah = newjadwal.Matakuliah
	if err := config.DB.Table("jadwal").Where("id", id).Save(&jadwal).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update matkul", jadwal))
}
