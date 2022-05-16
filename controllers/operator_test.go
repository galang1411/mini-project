package controllers

import (
	"bytes"
	"encoding/json"
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func InitEchoTestOperator() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}


// func GetToken() string {
// 	e := InitEchoTestOperator()
// 	InsertDataUserForGetOperator()
// 	operator := `{"username": "Alta", "password": "123"}`
// 	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(operator))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/login")
// 	LoginController(c)

// 	body := rec.Body.String()

// 	var response models.Token

// 	_ = json.Unmarshal([]byte(body), &response)

// 	return fmt.Sprintf("Bearer %s", response.Token)
// }

func TestPersonLogin(t *testing.T) {
	e := initTestEcho()
	operator1 := models.Operator{Username: "dono", Password: "rahasia"}
	operator1.ID = int(1)
	operator2 := models.Operator{Username: "kasino", Password: "rahasia"}
	operator2.ID = int(2)

	// login request
	login, err := json.Marshal(models.Operator{Username: "dono", Password: "rahasia"})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(login))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	// send request
	if err := LoginController(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	
	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}
	//
	// compare response
	var p models.Operator
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}

}
