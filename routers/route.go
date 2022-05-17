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
	eJwt.POST("/dosen", controllers.CreateDosencontrollers)
	eJwt.DELETE("/dosen/:nid", controllers.DeleteDosencontrollers)
	e.PUT("dosen/:nid", controllers.UpdateDosencontrollers)

	// Route to matkul handler function
	e.GET("/matkul", controllers.GetMatkulscontrollers)
	e.GET("/matkul/:id", controllers.GetMatkulcontrollers)
	eJwt.POST("/matkul", controllers.CreateMatkulcontrollers)
	eJwt.DELETE("/matkul/:id", controllers.DeleteMatkulcontrollers)
	e.PUT("matkul/:id", controllers.UpdateMatkulcontrollers)

	// Route to ruangan handler function
	e.GET("/ruangan", controllers.GetRuanganscontrollers)
	e.GET("/ruangan/:id", controllers.GetRuangancontrollers)
	eJwt.POST("/ruangan", controllers.CreateRuangancontrollers)
	eJwt.DELETE("/ruangan/:id", controllers.DeleteRuangancontrollers)
	e.PUT("ruangan/update/:id", controllers.UpdateRuangancontrollers)

	e.GET("/jadwal", controllers.GetJadwalscontrollers)
	e.GET("/jadwal/hari/:hari", controllers.GetHariJadwalHaricontrollers)
	e.GET("/jadwal/:id", controllers.GetJadwalIDcontrollers)
	eJwt.POST("/jadwal", controllers.CreateJadwalscontrollers)
	e.PUT("/jadwal/update/:id", controllers.UpdateJadwalcontrollers)
	eJwt.DELETE("/jadwal/:id", controllers.DeleteJadwalcontrollers)
	return e

	// go test ./controllers -coverprofile=coverage.out
}
