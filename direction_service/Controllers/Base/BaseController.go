package Base

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteDescription struct {
	Method string
	Path   string
}
type BaseController struct {
	Router gin.IRouter
	Routes []RouteDescription
}

func NewBaseController(router gin.IRouter) *BaseController {
	return &BaseController{Router: router}
}

func (controller BaseController) PrintRoutes() {
	//TODO implement me
	for i, route := range controller.Routes {
		fmt.Println(i+1, route.Method, route.Path)
	}
}

func (controller *BaseController) GET(path string, fn ...gin.HandlerFunc) gin.IRoutes {
	controller.Routes = append(controller.Routes, RouteDescription{
		Method: "GET",
		Path:   path,
	})
	return controller.Router.GET(path, fn...)
}

func (controller *BaseController) POST(path string, fn ...gin.HandlerFunc) gin.IRoutes {
	controller.Routes = append(controller.Routes, RouteDescription{
		Method: "POST",
		Path:   path,
	})
	return controller.Router.POST(path, fn...)
}

func (controller *BaseController) PUT(path string, fn ...gin.HandlerFunc) gin.IRoutes {
	controller.Routes = append(controller.Routes, RouteDescription{
		Method: "PUT",
		Path:   path,
	})
	return controller.Router.PUT(path, fn...)
}

func (controller *BaseController) DELETE(path string, fn ...gin.HandlerFunc) gin.IRoutes {
	controller.Routes = append(controller.Routes, RouteDescription{
		Method: "DELETE",
		Path:   path,
	})
	return controller.Router.DELETE(path, fn...)
}
