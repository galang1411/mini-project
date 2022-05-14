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

func InitEchoTestDosen() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

type DosenResponse struct {
	Message string
	Data    []models.Dosen
}

type ErrorResponse struct {
	Message string
}

func InsertDataUserForGetDosen() error {
	dosen := models.Dosen{
		NID:    67676868,
		Name:   "Alta",
		Gender: "L",
		Major:  "IT",
	}

	var err error
	if err = config.DB.Save(&dosen).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateDosenController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get dosen normal",
			path:       "/dosen",
			expectCode: http.StatusCreated,
		},
	}

	e := InitEchoTestDosen()
	dosen := `{"nid": 67676868, "nama": "Alta", "jenis_kelamin": "L", "jurusan" : "IT"}`
	req := httptest.NewRequest(http.MethodPost, "/dosen", strings.NewReader(dosen))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateDosencontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestGetAllDosenController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all dosen normal",
			path:       "/dosen",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestDosen()
	InsertDataUserForGetDosen()
	req := httptest.NewRequest(http.MethodGet, "/dosen", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, (GetDosenscontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// convert struct
			var dosens DosenResponse
			err := json.Unmarshal([]byte(body), &dosens)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(dosens.Data))
		}
	}
}

func TestGetDosenController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get dosen normal",
			path:       "/dosen/:nid",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestDosen()
	InsertDataUserForGetDosen()
	req := httptest.NewRequest(http.MethodGet, "/dosen/:nid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("nid")
		c.SetParamValues("67676868")

		if assert.NoError(t, (GetDosencontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestUpdateDosenController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update dosen normal",
			path:       "/dosen/:nid",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestDosen()
	InsertDataUserForGetDosen()
	dosen := `{"nid": 67676868, "nama": "Alta", "jenis_kelamin": "L", "jurusan" : "IT"}`
	req := httptest.NewRequest(http.MethodPut, "/jwt/users/:id", strings.NewReader(dosen))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("nid")
		c.SetParamValues("67676868")

		if assert.NoError(t, UpdateDosencontrollers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

func TestDeleteDosenController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete dosen normal",
			path:       "/dosen/:nid",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestDosen()
	InsertDataUserForGetDosen()
	req := httptest.NewRequest(http.MethodDelete, "/dosen/:nid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("nid")
		c.SetParamValues("67676868")

		if assert.NoError(t, (DeleteDosencontrollers)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}
