package router

import (
	"fmt"
	"github.com/gtxiqbal/altera-agmc/Day-2/config"
	"github.com/gtxiqbal/altera-agmc/Day-2/controller"
	"github.com/gtxiqbal/altera-agmc/Day-2/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RecoverWithConfig(config.RecoverConfig()))
	e.Use(middleware.RequestIDWithConfig(config.RequestIDConfig()))
	e.Use(middleware.LoggerWithConfig(config.LoggerConfig()))
	return &Router{e: e}
}

func (r *Router) StartServer() {
	r.e.Logger.Fatal(r.e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

func (r *Router) SetRouteUser(userService service.UserService) {
	userController := controller.NewUserController(userService)
	groupUser := r.e.Group("/users")
	groupUser.GET("", userController.GetAll)
	groupUser.GET("/:id", userController.GetByID)
	groupUser.POST("", userController.Create)
	groupUser.PUT("/:id", userController.Update)
	groupUser.DELETE("/:id", userController.Delete)
}

func (r *Router) SetRouteBook(bookService service.BookService) {
	bookController := controller.NewBookController(bookService)
	groupBook := r.e.Group("/books")
	groupBook.GET("", bookController.GetAll)
	groupBook.GET("/:id", bookController.GetByID)
	groupBook.POST("", bookController.Create)
	groupBook.PUT("/:id", bookController.Update)
	groupBook.DELETE("/:id", bookController.Delete)
}
