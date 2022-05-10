package routers

import (
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	// Route to Dosen handler function
	e.GET("/dosens", controllers.GetDosenscontrollers)
	e.GET("/dosen/:nid", controllers.GetDosencontrollers)
	e.POST("/dosen", controllers.CreateDosencontrollers)
	e.DELETE("/dosen/:nid", controllers.DeleteDosencontrollers)
	e.PUT("dosen/:nid", controllers.UpdateDosencontrollers)

	// Route to matkul handler function
	e.GET("/matkuls", controllers.GetMatkulscontrollers)
	e.GET("/matkul/:id", controllers.GetMatkulcontrollers)
	e.POST("/matkul", controllers.CreateMatkulcontrollers)
	e.DELETE("/matkul/:id", controllers.DeleteMatkulcontrollers)
	e.PUT("matkul/:id", controllers.UpdateMatkulcontrollers)

	// Route to ruangan handler function
	e.GET("/ruangans", controllers.GetRuanganscontrollers)
	e.GET("/ruangan/:id", controllers.GetRuangancontrollers)
	e.POST("/ruangan", controllers.CreateRuangancontrollers)
	e.DELETE("/ruangan/:id", controllers.DeleteRuangancontrollers)
	e.PUT("ruangan/:id", controllers.UpdateRuangancontrollers)

	return e
}
