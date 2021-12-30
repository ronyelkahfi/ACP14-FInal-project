package categories

import (
	"errors"
	_categoriesDomain "final-project/business/categories"
	_controllers "final-project/controllers"
	_request "final-project/controllers/categories/request"
	_response "final-project/controllers/categories/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	usecase _categoriesDomain.Usecase
}

func NewCategoryController(categoryUsecase _categoriesDomain.Usecase) *CategoryController {
	return &CategoryController{
		usecase: categoryUsecase,
	}
}

func (controller *CategoryController) GetCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := controller.usecase.GetCategory(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(categories))
}

func (controller *CategoryController) CreateCategory(c echo.Context) error {
	var data _request.CategoryRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	dataDomain := _categoriesDomain.Domain{
		Name: data.Name,
	}
	_, error := controller.usecase.CreateCategory(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "")
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	affectedRow, error := controller.usecase.DeleteCategory(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}

	if affectedRow == 0 {
		return _controllers.NewErrorResponse(c, errors.New("Id not found in database"))
	}
	return _controllers.NewSuccessResponse(c, "")
}
