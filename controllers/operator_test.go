package controllers

import (
	"encoding/json"
	"fmt"
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestOperator() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type OperatorResponse struct {
	Message string
	Data    []models.Operator
}

func GetToken() string {
	e := InitEchoTestOperator()
	InsertDataUserForGetOperator()
	operator := `{"username": "Alta", "password": "123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(operator))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	LoginController(c)

	body := rec.Body.String()

	var response models.Token

	_ = json.Unmarshal([]byte(body), &response)

	return fmt.Sprintf("Bearer %s", response.Token)
}

func InsertDataUserForGetOperator() error {
	operator := models.Operator{
		Username: "Alta",
		Password: "123",
	}

	var err error
	if err = config.DB.Save(&operator).Error; err != nil {
		return err
	}
	return nil
}

func LoginOperatorController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "login normal",
			path:       "/login",
			expectCode: http.StatusOK,
		},
	}
	e := InitEchoTestOperator()
	InsertDataUserForGetOperator()
	operator := `{"username": "Alta", "password": "123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(operator))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, LoginController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
