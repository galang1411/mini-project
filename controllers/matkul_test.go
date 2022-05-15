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

func InitEchoTestMatkul() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type MatkulResponse struct {
	Message string
	Data    []models.Matakuliah
}

func InsertDataUserForGetMatkul() error {
	Matkul := models.Matakuliah{
		Name:     "Alta",
		SKS:      3,
		Semester: 3,
	}

	var err error
	if err = config.DB.Save(&Matkul).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateMatkulController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get matkul normal",
			path:       "/matkul",
			expectCode: http.StatusCreated,
		},
	}

	e := InitEchoTestMatkul()
	Matkul := `{"nama": "Alta", "sks": 3, "semester" : 3}`
	req := httptest.NewRequest(http.MethodPost, "/matkul", strings.NewReader(Matkul))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateMatkulcontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestGetAllMatkulController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all Matkul normal",
			path:       "/matkul",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestMatkul()
	InsertDataUserForGetMatkul()
	req := httptest.NewRequest(http.MethodGet, "/matkul", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, (GetMatkulscontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// convert struct
			var Matkuls MatkulResponse
			err := json.Unmarshal([]byte(body), &Matkuls)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(Matkuls.Data))
		}
	}
}

func TestGetMatkulController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get Matkul normal",
			path:       "/matkul/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestMatkul()
	InsertDataUserForGetMatkul()
	req := httptest.NewRequest(http.MethodGet, "/matkul/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, (GetMatkulcontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateMatkulController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update Matkul normal",
			path:       "/matkul/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestMatkul()
	InsertDataUserForGetMatkul()
	Matkul := `{"nama": "Alta", "sks": 3, "semester" : 3}`
	req := httptest.NewRequest(http.MethodPost, "/matkul", strings.NewReader(Matkul))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateMatkulcontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteMatkulController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete Matkul normal",
			path:       "/matkul/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestMatkul()
	InsertDataUserForGetMatkul()
	req := httptest.NewRequest(http.MethodDelete, "/matkul/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, (DeleteMatkulcontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
