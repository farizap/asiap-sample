package http

import (
	common "asiap/pkg/common/http"
	"asiap/pkg/user/business"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get item API controller
type Controller struct {
	service   business.UserService
	validator *v10.Validate
}

//NewController Construct item API controller
func NewController(service business.UserService) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) GetUserRegistrationByManagerID(c echo.Context) error {
	ID := c.Param("id")
	item, err := controller.service.UserByManagerID(ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := NewUsersResponse(item)
	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) ApproveUserRegistration(c echo.Context) error {
	ID := c.Param("id")

	err := controller.service.ApproveRegistration(ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := NewAddUserRegistrationResponse(ID)
	return c.JSON(http.StatusCreated, response)
}

func (controller *Controller) AddUserRegistration(c echo.Context) error {
	createItemRequest := new(AddUserRegistration)

	if err := c.Bind(createItemRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	cmd := business.AddUserRegistrationCommand{ID: createItemRequest.ID, Name: createItemRequest.Name, Email: createItemRequest.Email, Location: createItemRequest.Location, ManagerID: createItemRequest.ManagerID}
	err := controller.service.AddUserRegistration(cmd)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	response := NewAddUserRegistrationResponse(cmd.ID)
	return c.JSON(http.StatusCreated, response)
}
