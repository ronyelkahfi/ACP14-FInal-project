package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("apps/configs/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// configDB := _dbDriver.ConfigDB{
	// 	DB_Username: viper.GetString(`database.user`),
	// 	DB_Password: viper.GetString(`database.pass`),
	// 	DB_Host:     viper.GetString(`database.host`),
	// 	DB_Port:     viper.GetString(`database.port`),
	// 	DB_Database: viper.GetString(`database.name`),
	// }

	// db := configDB.ConnectDB()
	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// userRepo := _userRepo.NewUserRepository(db)
	// userUseCase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	// userController := _userController.NewUserController(userUseCase)

	// productRepo := _productRepo.NewProductRepository(db)
	// productUseCase := _productUsecase.NewProductUsecase(productRepo, timeoutContext)
	// productController := _productController.NewProductController(productUseCase)

	// categoryRepo := _categoryRepo.NewCategoryRepository(db)
	// categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepo, timeoutContext)
	// categoryController := _categoryController.NewCategoryController(categoryUseCase)

	// transactionRepo := _transactionRepo.NewTransactionRepository(db)
	// transactionDetailRepo := _transactionDetailRepo.NewTransDetailRepository(db)
	// transactionUseCase := _transactionUsecase.NewTransactionUsecase(transactionRepo, transactionDetailRepo, timeoutContext)
	// transactionController := _transactionController.NewTransactionController(transactionUseCase)

	e := echo.New()
	// routesInit := _route.ControllerList{
	// 	UserController:        *userController,
	// 	ProductController:     *productController,
	// 	CategoryController:    *categoryController,
	// 	TransactionController: *transactionController,
	// }
	// routesInit.RouteRegister(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})
	log.Fatal(e.Start(viper.GetString("server.address")))
}
