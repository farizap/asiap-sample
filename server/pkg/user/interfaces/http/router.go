package http

import (
	"github.com/labstack/echo"
)

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, userController *Controller) {
	if userController == nil {
		panic("item controller cannot be nil")
	}

	//item
	itemV1 := e.Group("v1/user")
	itemV1.GET("/bymanager/:id", userController.GetUserRegistrationByManagerID)
	itemV1.POST("/:id/approve", userController.ApproveUserRegistration)
	itemV1.POST("/request", userController.AddUserRegistration)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
