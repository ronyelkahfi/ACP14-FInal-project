package routes

import (
	_categories "final-project/controllers/categories"
	_products "final-project/controllers/products"
	_transactions "final-project/controllers/transactions"
	_users "final-project/controllers/users"

	// _jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController        _users.UserController
	ProductController     _products.ProductController
	CategoryController    _categories.CategoryController
	TransactionController _transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	r := c.Group("/jwt")
	r.Use(middleware.JWT([]byte("ThisisSecretGais")))

	c.GET("/users", cl.UserController.GetUser)
	c.POST("/users", cl.UserController.Register)
	c.GET("/products", cl.ProductController.GetProduct)
	c.POST("/products", cl.ProductController.CreateProduct)
	c.DELETE("/products/:id", cl.ProductController.DeleteProduct)
	c.GET("/categories", cl.CategoryController.GetCategory)
	c.POST("/categories", cl.CategoryController.CreateCategory)
	c.DELETE("/categories/:id", cl.CategoryController.DeleteCategory)
	c.POST("/login", cl.UserController.Login)
	r.POST("/orders", cl.TransactionController.CreateTransaction)
	r.DELETE("/orders/:id", cl.TransactionController.DeleteTransaction)
	r.DELETE("/orders/detail/:id", cl.TransactionController.DeleteDetailTransaction)
	c.POST("/register", cl.UserController.Register)
}
