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
	e.PUT("dosen/:nid", controllers.UpdateDosencontrollers)
	eJwt.POST("/dosen", controllers.CreateDosencontrollers)
	eJwt.DELETE("/dosen/:nid", controllers.DeleteDosencontrollers)

	// Route to matkul handler function
	e.GET("/matkul", controllers.GetMatkulscontrollers)
	e.GET("/matkul/:id", controllers.GetMatkulcontrollers)
	e.PUT("matkul/:id", controllers.UpdateMatkulcontrollers)
	eJwt.POST("/matkul", controllers.CreateMatkulcontrollers)
	eJwt.DELETE("/matkul/:id", controllers.DeleteMatkulcontrollers)

	// Route to ruangan handler function
	e.GET("/ruangan", controllers.GetRuanganscontrollers)
	e.GET("/ruangan/:id", controllers.GetRuangancontrollers)
	e.PUT("ruangan/update/:id", controllers.UpdateRuangancontrollers)
	eJwt.POST("/ruangan", controllers.CreateRuangancontrollers)
	eJwt.DELETE("/ruangan/:id", controllers.DeleteRuangancontrollers)

	e.GET("/jadwal", controllers.GetJadwalscontrollers)
	e.GET("/jadwal/hari/:hari", controllers.GetHariJadwalHaricontrollers)
	e.GET("/jadwal/:id", controllers.GetJadwalIDcontrollers)
	e.PUT("/jadwal/update/:id", controllers.UpdateJadwalcontrollers)
	eJwt.POST("/jadwal", controllers.CreateJadwalscontrollers)
	eJwt.DELETE("/jadwal/:id", controllers.DeleteJadwalcontrollers)
	return e

	// go test ./controllers -coverprofile=coverage.out
}
