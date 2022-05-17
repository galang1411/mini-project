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

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new ruangan", ruangan))
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
	id, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	fmt.Println("Isi userkId ", id)
	var r models.Ruangan
	fmt.Printf("Isi userk sebelum select %#v\n", r)
	if err := config.DB.First(&r, id).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if r.ID == 0 {
		return c.String(http.StatusNotFound, "userk not found")
	}
	fmt.Printf("Isi userk setelah select %#v\n", r)
	if err := c.Bind(&r); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	fmt.Printf("Isi user setelah bind %#v\n", r)
	fmt.Printf("Before update: %#v\n", r)
	if err := config.DB.Save(&r).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, r)

	// id, _ := strconv.Atoi(c.Param("id"))
	// ruangan := models.Ruangan{}
	// if err := config.DB.Table("ruangan").First(&ruangan, id).Error; err != nil {
	// 	if err.Error() == "record not found" {
	// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
	// 			"message": "ruangan not found",
	// 		})
	// 	}

	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	// newruangan := models.Ruangan{}
	// c.Bind(&newruangan)

	// ruangan.Name = newruangan.Name

	// if err := config.DB.Table("ruangan").Save(&ruangan).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// return c.JSON(http.StatusOK, helper.BuildResponse("success update ruangan", ruangan))
}
