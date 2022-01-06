package main

import (
	_categoryUsecase "final-project/business/categories"
	_productUsecase "final-project/business/products"
	_transactionUsecase "final-project/business/transactions"
	_userUsecase "final-project/business/users"
	_dbDriver "final-project/drivers/databases"
	_categoryRepo "final-project/drivers/databases/categories"
	_productRepo "final-project/drivers/databases/products"
	_transactionDetailRepo "final-project/drivers/databases/transactionDetails"
	_transactionRepo "final-project/drivers/databases/transactions"
	_userRepo "final-project/drivers/databases/users"
	"log"
	"os"
	"time"

	_route "final-project/apps/routes"
	_categoryController "final-project/controllers/categories"
	_productController "final-project/controllers/products"
	_transactionController "final-project/controllers/transactions"
	_userController "final-project/controllers/users"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load .env file")
	}

	configDB := _dbDriver.ConfigDB{
		DB_Username: os.Getenv("DBUSERNAME"),
		DB_Password: os.Getenv("DBPASSWORD"),
		DB_Host:     os.Getenv("DBHOST"),
		DB_Port:     os.Getenv("DBPORT"),
		DB_Database: os.Getenv("DBNAME"),
	}
	// fmt.Println(os.Getenv("DBUSERNAME"), os.Getenv("DBPASSWORD"))
	db := configDB.ConnectDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _userRepo.NewUserRepository(db)
	userUseCase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	productRepo := _productRepo.NewProductRepository(db)
	productUseCase := _productUsecase.NewProductUsecase(productRepo, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepo, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	transactionRepo := _transactionRepo.NewTransactionRepository(db)
	transactionDetailRepo := _transactionDetailRepo.NewTransDetailRepository(db)
	transactionUseCase := _transactionUsecase.NewTransactionUsecase(transactionRepo, transactionDetailRepo, timeoutContext)
	transactionController := _transactionController.NewTransactionController(transactionUseCase)

	e := echo.New()
	routesInit := _route.ControllerList{
		UserController:        *userController,
		ProductController:     *productController,
		CategoryController:    *categoryController,
		TransactionController: *transactionController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
