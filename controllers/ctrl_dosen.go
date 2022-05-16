package controllers

import (
	"fmt"
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all dosen
func GetDosenscontrollers(c echo.Context) error {
	var dosen []models.Dosen
	if err := config.DB.Find(&dosen).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all dosens", dosen))
}

// get dosen by nid
func GetDosencontrollers(c echo.Context) error {
	nid, _ := strconv.Atoi(c.Param("nid"))
	dosen := models.Dosen{}
	if err := config.DB.First(&dosen, nid).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "dosen not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get dosen", dosen))
}

// create dosen by nid
func CreateDosencontrollers(c echo.Context) error {
	dosen := models.Dosen{}
	c.Bind(&dosen)

	if err := config.DB.Debug().Create(&dosen).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new dosen", dosen))
}

// delete dosen by nid
func DeleteDosencontrollers(c echo.Context) error {
	nid, _ := strconv.Atoi(c.Param("nid"))
	dosen := models.Dosen{}
	if err := config.DB.Where("nid = ?", nid).First(&dosen).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "dosen not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("nid = ?", nid).Delete(&dosen).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("dosen deleted successfully", dosen))
}

// update dosen by nid
func UpdateDosencontrollers(c echo.Context) error {
	nid := c.Param("nid")
	dosen := models.Dosen{}

	if err := config.DB.First(&dosen, "nid = ?", nid).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "dosen not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newdosen := models.Dosen{}
	c.Bind(&newdosen)
	fmt.Println("dosen", dosen)
	dosen.Name = newdosen.Name
	dosen.Gender = newdosen.Gender
	dosen.Major = newdosen.Major
	if err := config.DB.Where("nid = ?", nid).Save(&dosen).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update dosen", dosen))
}
