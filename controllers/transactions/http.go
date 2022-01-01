package transactions

import (
	"errors"
	middleware "final-project/apps/middlewares"
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

func (controller *TransactionController) CreateTransaction(c echo.Context) error {
	// mdw.ExtractTokenMetadata(c)

	userId := middleware.ExtractToken(c)
	if userId == 0 {
		return _controllers.NewErrorResponse(c, errors.New("Token Error"))
	}
	var data _request.TransactionRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	transactions, err := controller.usecase.NewTransaction(ctx, userId, _request.DetailToDomain(data))

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
