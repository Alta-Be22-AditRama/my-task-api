package routers

import (
	// "be22/clean-arch/features/product/data"
	// "be22/clean-arch/features/product/service"
	// "be22/clean-arch/features/product/handler"
	"my-task-api/app/middlewares"
	_userData "my-task-api/features/user/data"
	_userHandler "my-task-api/features/user/handler"
	_userService "my-task-api/features/user/service"
	"my-task-api/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	// factory
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	// productData := data.New(db)
	// productService := service.New(productData)
	// productHandlerAPI := handler.New(productService)

	// e.GET("/users", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "hello world",
	// 	})
	// })
	e.POST("/users", userHandlerAPI.Register) //register
	e.POST("/login", userHandlerAPI.Login) //login
	e.GET("/users", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.PUT("/users",userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	
	
	// e.GET("/product", productHandlerAPI.GetAll)
}
