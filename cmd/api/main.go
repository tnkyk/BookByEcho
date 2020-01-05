package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/handler/rest"
	"github.com/tnkyk/BookByEcho/infra/persistence"
	"github.com/tnkyk/BookByEcho/usecase"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := rest.NewUserHandler(userUseCase)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/all", func(c echo.Context) error { return userHandler.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userHandler.DeleteUser(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userHandler.DeleteUser(c) })
	e.PUT("/user/:id/", func(c echo.Context) error { return userHandler.UpdateUser(c) })
	e.POST("/user/signup/", func(c echo.Context) error { return userHandler.SignUp(c) })
	e.POST("/user/login", func(c echo.Context) error { return userHandler.SignIn(c) })

	e.GET("/book/all", Bookhandler.GetBooks)
	e.GET("/book/:id", Bookhandler.GetBookByID)
	e.DELETE("/book/:id", Bookhandler.DeleteBook)
	e.POST("/book/Register", Bookhandler.Register)
	e.POST("/book/:id", Bookhandler.Reservation)

	e.Logger.Fatal(e.Start(":1323"))
}
