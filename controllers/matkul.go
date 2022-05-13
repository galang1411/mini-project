package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all matkuls
func GetMatkulscontrollers(c echo.Context) error {
	var matkuls []models.Matakuliah
	if err := config.DB.Find(&matkuls).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all matakuliah", matkuls))
}

// get matkul by id
func GetMatkulcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	matkul := models.Matakuliah{}
	if err := config.DB.First(&matkul, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "matkul not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get matakuliah", matkul))
}

// create matkul by id
func CreateMatkulcontrollers(c echo.Context) error {
	matkul := models.Matakuliah{}
	c.Bind(&matkul)

	if err := config.DB.Save(&matkul).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new matakuliah", matkul))
}

// delete matkul by id
func DeleteMatkulcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	matkul := models.Matakuliah{}
	if err := config.DB.Table("matakuliah").First(&matkul, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "matkul not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&matkul).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("matakuliah deleted successfully", matkul))
}

// update matkul by id
func UpdateMatkulcontrollers(c echo.Context) error {
	id := c.Param("id")
	matkul := models.Matakuliah{}

	if err := config.DB.First(&matkul, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "matkul not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newmatkul := models.Matakuliah{}
	c.Bind(&newmatkul)

	matkul.Name = newmatkul.Name
	matkul.SKS = newmatkul.SKS
	matkul.Semester = newmatkul.Semester
	if err := config.DB.Save(&matkul).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update matkul", matkul))
}
