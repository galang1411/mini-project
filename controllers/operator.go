package controllers

import (
	"mini-project/config"
	"mini-project/middleware"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	operator := models.Operator{}
	c.Bind(&operator)

	if err := config.DB.Where("username = ? AND password = ?", operator.Username, operator.Password).First(&operator).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.CreateToken(operator.ID, operator.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "berhasil login",
		"operator": token,
	})
}
