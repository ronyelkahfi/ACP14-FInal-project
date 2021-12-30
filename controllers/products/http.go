package products

import (
	_productsDomain "final-project/business/products"
	_controllers "final-project/controllers"
	_request "final-project/controllers/products/request"
	_response "final-project/controllers/products/response"
	"final-project/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	usecase _productsDomain.Usecase
}

func NewProductController(productUsecase _productsDomain.Usecase) *ProductController {
	return &ProductController{
		usecase: productUsecase,
	}
}

func (controller *ProductController) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := controller.usecase.GetProduct(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) CreateProduct(c echo.Context) error {
	var data _request.ProductRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	dataDomain := _productsDomain.Domain{
		Name:       data.Name,
		Price:      uint(data.Price),
		CategoryId: uint(data.CategoryId),
	}
	_, error := controller.usecase.CreateProduct(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "")
}

func (controller *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	affectedRow, error := controller.usecase.DeleteProduct(ctx, id)
	if affectedRow == 0 {
		return _controllers.NewErrorResponse(c, helpers.ErrIdNotFound)
	}
	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "")
}
