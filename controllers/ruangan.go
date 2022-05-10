package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all ruangans
func GetRuanganscontrollers(c echo.Context) error {
	var ruangan []models.Ruangan
	if err := config.DB.Find(&ruangan).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all ruangan", ruangan))
}

// get ruangan by id
func GetRuangancontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ruangan := models.Ruangan{}
	if err := config.DB.First(&ruangan, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "ruangan not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get ruangan", ruangan))
}

// create ruangan by id
func CreateRuangancontrollers(c echo.Context) error {
	ruangan := models.Ruangan{}
	c.Bind(&ruangan)

	if err := config.DB.Save(&ruangan).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new user", ruangan))
}

// delete ruangan by id
func DeleteRuangancontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ruangan := models.Ruangan{}
	if err := config.DB.Table("ruangan").First(&ruangan, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "ruangan not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&ruangan).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("ruangan deleted successfully", ruangan))
}

// update ruangan by id
func UpdateRuangancontrollers(c echo.Context) error {
	id := c.Param("id")
	ruangan := models.Ruangan{}

	if err := config.DB.First(&ruangan, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "ruangan not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newruangan := models.Ruangan{}
	c.Bind(&newruangan)

	ruangan.Name = newruangan.Name

	if err := config.DB.Save(&ruangan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update ruangan", ruangan))
}
