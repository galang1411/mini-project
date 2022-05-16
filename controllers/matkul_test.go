package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func initTestEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

func TestMatkulAdd(t *testing.T) {
	e := initTestEcho()

	// compose request
	newMatkul, err := json.Marshal(map[string]interface{}{
		"nama":     "Alterra",
		"sks":      3,
		"semester": 4,
	})

	fmt.Println(string(newMatkul))

	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newMatkul))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/matkul")

	// send request
	if err = CreateMatkulcontrollers(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 201 {
		t.Errorf("should return 201, get: %d", rec.Code)
	}

	// compare response
	m := models.Matakuliah{}
	if err = json.Unmarshal(rec.Body.Bytes(), &m); err != nil {
		t.Errorf("unmarshalling returned failed")
	}
	fmt.Println(m)
	expectedNama := "Alterra"
	if m.Name != expectedNama {
		t.Errorf("name should be %s, get: %s", expectedNama, m.Name)
	}
	expectedSKS := 3
	if m.SKS != expectedSKS {
		t.Errorf("sks should be %d, get: %d", expectedSKS, m.SKS)
	}
	expectedSemester := 4
	if m.Semester != expectedSemester {
		t.Errorf("semester should be %d, get: %d", expectedSemester, m.Semester)
	}
}

// func TestGetAll(t *testing.T) {
// 	e := initTestEcho()
// 	matkul1 := models.Matakuliah{Name: "alterra"}
// 	matkul1.ID = uint(1)
// 	matkul2 := models.Matakuliah{Name: "alta"}
// 	matkul2.ID = uint(2)

// 	// get all request
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)

// 	req.Header.Set("Content-Type", "application/json")
// 	rec := httptest.NewRecorder()
// 	context := e.NewContext(req, rec)
// 	context.SetPath("/matkul")

// 	// get all
// 	if err := GetMatkulscontrollers(context); err != nil {
// 		t.Errorf("should not get error, get error: %s", err)
// 		return
// 	}

// 	// compare status
// 	if rec.Code != 200 {
// 		t.Errorf("should return 200, get: %d", rec.Code)
// 	}

// 	var pList []models.Matakuliah
// 	if err := json.Unmarshal(rec.Body.Bytes(), &pList); err != nil {
// 		t.Errorf("unmarshalling returned person list failed")
// 	}

// 	expectedPListLength := 2
// 	if len(pList) != expectedPListLength {
// 		t.Errorf("expecting pList's length to be %d, get %d, data: %#v", expectedPListLength, len(pList), pList)
// 	}
// }

// func InitEchoTestMatkul() *echo.Echo {
// 	config.InitDBTest()
// 	e := echo.New()
// 	return e
// }

// type MatkulResponse struct {
// 	Message string
// 	Data    []models.Matakuliah
// }

// func InsertDataUserForGetMatkul() error {
// 	Matkul := models.Matakuliah{
// 		Name:     "Alta",
// 		SKS:      3,
// 		Semester: 3,
// 	}

// 	var err error
// 	if err = config.DB.Save(&Matkul).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func TestCreateMatkulController(t *testing.T) {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 	}{
// 		{
// 			name:       "get matkul normal",
// 			path:       "/matkul",
// 			expectCode: http.StatusCreated,
// 		},
// 	}

// 	e := InitEchoTestMatkul()
// 	Matkul := `{"nama": "Alta", "sks": 3, "semester" : 3}`
// 	req := httptest.NewRequest(http.MethodPost, "/matkul", strings.NewReader(Matkul))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, CreateMatkulcontrollers(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 		}
// 	}
// }

// func TestGetAllMatkulController(t *testing.T) {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 		sizeData   int
// 	}{
// 		{
// 			name:       "get all Matkul normal",
// 			path:       "/matkul",
// 			expectCode: http.StatusOK,
// 			sizeData:   1,
// 		},
// 	}

// 	e := InitEchoTestMatkul()
// 	InsertDataUserForGetMatkul()
// 	req := httptest.NewRequest(http.MethodGet, "/matkul", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, (GetMatkulscontrollers)(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 			body := rec.Body.String()

// 			// convert struct
// 			var Matkuls MatkulResponse
// 			err := json.Unmarshal([]byte(body), &Matkuls)

// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, testCase.sizeData, len(Matkuls.Data))
// 		}
// 	}
// }

// func TestGetMatkulController(t *testing.T) {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 		sizeData   int
// 	}{
// 		{
// 			name:       "get Matkul normal",
// 			path:       "/matkul/:id",
// 			expectCode: http.StatusOK,
// 		},
// 	}

// 	e := InitEchoTestMatkul()
// 	InsertDataUserForGetMatkul()
// 	req := httptest.NewRequest(http.MethodGet, "/matkul/:id", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")

// 		if assert.NoError(t, (GetMatkulcontrollers)(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 		}
// 	}
// }

// func TestUpdateMatkulController(t *testing.T) {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 	}{
// 		{
// 			name:       "update Matkul normal",
// 			path:       "/matkul/:id",
// 			expectCode: http.StatusOK,
// 		},
// 	}

// 	e := InitEchoTestMatkul()
// 	InsertDataUserForGetMatkul()
// 	Matkul := `{"nama": "Alta", "sks": 3, "semester" : 3}`
// 	req := httptest.NewRequest(http.MethodPost, "/matkul", strings.NewReader(Matkul))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")

// 		if assert.NoError(t, UpdateMatkulcontrollers(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 		}
// 	}
// }

// func TestDeleteMatkulController(t *testing.T) {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 	}{
// 		{
// 			name:       "delete Matkul normal",
// 			path:       "/matkul/:id",
// 			expectCode: http.StatusOK,
// 		},
// 	}

// 	e := InitEchoTestMatkul()
// 	InsertDataUserForGetMatkul()
// 	req := httptest.NewRequest(http.MethodDelete, "/matkul/:id", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")

// 		if assert.NoError(t, (DeleteMatkulcontrollers)(c)) {
// 			assert.Equal(t, testCase.expectCode, rec.Code)
// 		}
// 	}
// }
