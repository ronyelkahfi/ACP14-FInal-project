package routes

import (
	_categories "final-project/controllers/categories"
	_products "final-project/controllers/products"
	_users "final-project/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController     _users.UserController
	ProductController  _products.ProductController
	CategoryController _categories.CategoryController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetUser)
	c.POST("/users", cl.UserController.Register)
	c.GET("/products", cl.ProductController.GetProduct)
	c.POST("/products", cl.ProductController.CreateProduct)
	c.DELETE("/products/:id", cl.ProductController.DeleteProduct)
	c.GET("/categories", cl.CategoryController.GetCategory)
	c.POST("/categories", cl.CategoryController.CreateCategory)
	c.DELETE("/categories/:id", cl.CategoryController.DeleteCategory)
}
