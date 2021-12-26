package routes

import (
	_products "final-project/controllers/products"
	_users "final-project/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController    _users.UserController
	ProductController _products.ProductController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetUser)
	c.POST("/users", cl.UserController.Register)
	c.GET("/products", cl.ProductController.GetProduct)
	c.POST("/products", cl.ProductController.CreateProduct)
	c.DELETE("/products/:id", cl.ProductController.DeleteProduct)
}
