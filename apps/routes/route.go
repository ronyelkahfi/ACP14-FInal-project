package routes

import (
	_categories "final-project/controllers/categories"
	_products "final-project/controllers/products"
	_transactions "final-project/controllers/transactions"
	_users "final-project/controllers/users"
	"net/http"

	_middlewares "final-project/apps/middlewares"

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
	// Middleware
	c.Use(middleware.Logger())
	c.Use(middleware.Recover())

	c.GET("/users", cl.UserController.GetUser)
	c.POST("/users", cl.UserController.Register)
	c.GET("/products", cl.ProductController.GetProduct)
	c.POST("/products", cl.ProductController.CreateProduct)
	c.DELETE("/products/:id", cl.ProductController.DeleteProduct)
	c.GET("/categories", cl.CategoryController.GetCategory)
	c.POST("/categories", cl.CategoryController.CreateCategory)
	c.DELETE("/categories/:id", cl.CategoryController.DeleteCategory)
	c.POST("/login", cl.UserController.Login)
	c.POST("/register", cl.UserController.Register)
	// Restricted group
	r := c.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &_middlewares.JwtCustomClaims{},
		SigningKey: []byte("SECRETges%&*^&*^*&("),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("/transactions", func(c echo.Context) error {
		claims, _ := _middlewares.ExtractClaims(c)
		return c.String(http.StatusOK, claims.Name)
	})
	r.POST("/orders", cl.TransactionController.CreateTransaction)
	r.DELETE("/orders/:id", cl.TransactionController.DeleteTransaction)
	r.DELETE("/orders/detail/:id", cl.TransactionController.DeleteDetailTransaction)
}
