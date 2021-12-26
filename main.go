package main

import (
	_categoryUsecase "final-project/business/categories"
	_productUsecase "final-project/business/products"
	_userUsecase "final-project/business/users"
	_dbDriver "final-project/drivers/databases"
	_categoryRepo "final-project/drivers/databases/categories"
	_productRepo "final-project/drivers/databases/products"
	_userRepo "final-project/drivers/databases/users"
	"log"
	"time"

	_route "final-project/apps/routes"
	_categoryController "final-project/controllers/categories"
	_productController "final-project/controllers/products"
	_userController "final-project/controllers/users"

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
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

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

	e := echo.New()
	routesInit := _route.ControllerList{
		UserController:     *userController,
		ProductController:  *productController,
		CategoryController: *categoryController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
