package users

import (
	_usersDomain "final-project/business/users"
	_controllers "final-project/controllers"
	_request "final-project/controllers/users/request"
	_response "final-project/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase _usersDomain.Usecase
}

func NewUserController(userUsecase _usersDomain.Usecase) *UserController {
	return &UserController{
		usecase: userUsecase,
	}
}

func (controller *UserController) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := controller.usecase.GetUser(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(users))
}

type Usercontoh struct {
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}

func (controller *UserController) Register(c echo.Context) error {
	var data _request.UserRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return err
	}
	dataDomain := _usersDomain.Domain{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	_, error := controller.usecase.Register(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "sss")
}
