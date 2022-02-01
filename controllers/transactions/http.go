package transactions

import (
	// middleware "final-project/apps/middlewares"
	_TransactionsDomain "final-project/business/transactions"
	_controllers "final-project/controllers"
	_request "final-project/controllers/transactions/request"
	_response "final-project/controllers/transactions/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	usecase _TransactionsDomain.Usecase
}

func NewTransactionController(TransactionUsecase _TransactionsDomain.Usecase) *TransactionController {
	return &TransactionController{
		usecase: TransactionUsecase,
	}
}

// type claims struct {
// 	userid
// }

func (controller *TransactionController) CreateTransaction(c echo.Context) error {

	// jwtclaims, authenticate := _middleware.ExtractClaims()
	// if authenticate == false {
	// 	return _controllers.NewErrorResponse(c, errors.New("Token error"))
	// }

	var data _request.TransactionRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	transactions, err := controller.usecase.NewTransaction(ctx, 1, _request.DetailToDomain(data))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(transactions))
}

func (controller *TransactionController) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, error := controller.usecase.DeleteTransaction(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "")
}
func (controller *TransactionController) DeleteDetailTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, error := controller.usecase.DeleteDetailTransaction(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "")
}
