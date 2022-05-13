package controllers

import (
	"encoding/json"
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestRuangan() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

type RuanganResponse struct {
	Message string
	Data    []models.Ruangan
}

func InsertDataUserForGetRuangan() error {
	Ruangan := models.Ruangan{
		Name: "Alta",
	}

	var err error
	if err = config.DB.Save(&Ruangan).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateRuanganController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get Ruangan normal",
			path:       "/ruangan",
			expectCode: http.StatusCreated,
		},
	}

	e := InitEchoTestRuangan()
	Ruangan := `{"nama": "Alta"}`
	req := httptest.NewRequest(http.MethodPost, "/Ruangan", strings.NewReader(Ruangan))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateRuangancontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestGetAllRuanganController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all Ruangan normal",
			path:       "/ruangan",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestRuangan()
	InsertDataUserForGetRuangan()
	req := httptest.NewRequest(http.MethodGet, "/ruangan", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, (GetRuanganscontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// convert struct
			var Ruangans RuanganResponse
			err := json.Unmarshal([]byte(body), &Ruangans)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(Ruangans.Data))
		}
	}
}

func TestGetRuanganController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get Ruangan normal",
			path:       "/ruangan/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestRuangan()
	InsertDataUserForGetRuangan()
	req := httptest.NewRequest(http.MethodGet, "/ruangan/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, (GetRuangancontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateRuanganController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update Ruangan normal",
			path:       "/ruangan/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestRuangan()
	InsertDataUserForGetRuangan()
	Ruangan := `{"nama": "Alta"}`
	req := httptest.NewRequest(http.MethodPost, "/ruangan", strings.NewReader(Ruangan))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateRuangancontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteRuanganController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete Ruangan normal",
			path:       "/ruangan/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestRuangan()
	InsertDataUserForGetRuangan()
	req := httptest.NewRequest(http.MethodDelete, "/ruangan/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, (DeleteRuangancontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
