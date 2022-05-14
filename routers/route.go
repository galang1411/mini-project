package routers

import (
	"mini-project/constants"
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	e.POST("/login", controllers.LoginController)
	// Route to Dosen handler function
	e.GET("/dosen", controllers.GetDosenscontrollers)
	e.GET("/dosen/:nid", controllers.GetDosencontrollers)
	e.POST("/dosen", controllers.CreateDosencontrollers)
	e.DELETE("/dosen/:nid", controllers.DeleteDosencontrollers)
	e.PUT("dosen/:nid", controllers.UpdateDosencontrollers)

	// Route to matkul handler function
	e.GET("/matkul", controllers.GetMatkulscontrollers)
	e.GET("/matkul/:id", controllers.GetMatkulcontrollers)
	e.POST("/matkul", controllers.CreateMatkulcontrollers)
	e.DELETE("/matkul/:id", controllers.DeleteMatkulcontrollers)
	e.PUT("matkul/:id", controllers.UpdateMatkulcontrollers)

	// Route to ruangan handler function
	e.GET("/ruangan", controllers.GetRuanganscontrollers)
	e.GET("/ruangan/:id", controllers.GetRuangancontrollers)
	e.POST("/ruangan", controllers.CreateRuangancontrollers)
	e.DELETE("/ruangan/:id", controllers.DeleteRuangancontrollers)
	e.PUT("ruangan/:id", controllers.UpdateRuangancontrollers)

	e.GET("/jadwal", controllers.GetJadwalscontrollers)
	e.GET("/jadwal/:hari", controllers.GetHariJadwalHaricontrollers)
	e.GET("/jadwal/:id", controllers.GetJadwalIDcontrollers)
	e.POST("/jadwal", controllers.CreateJadwalscontrollers)
	e.PUT("/jadwal/:id", controllers.UpdateJadwalcontrollers)
	e.DELETE("/jadwal/:id", controllers.DeleteJadwalcontrollers)
	return e

	// go test ./controllers -coverprofile=coverage.out
}
