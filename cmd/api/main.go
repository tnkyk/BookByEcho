package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tnkyk/BookByEcho/handler/rest"
	"github.com/tnkyk/BookByEcho/infra/persistence"
	"github.com/tnkyk/BookByEcho/usecase"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := rest.NewUserHandler(userUseCase)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := middleware.JWTConfig{
		Claims:     &rest.Claims{},
		SigningKey: []byte("my_secret_key"),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/all", func(c echo.Context) error { return userHandler.Index(c) }, middleware.JWTWithConfig(config))
	e.GET("/user/:id", func(c echo.Context) error { return userHandler.DeleteUser(c) }, middleware.JWTWithConfig(config))
	e.DELETE("/user/:id", func(c echo.Context) error { return userHandler.DeleteUser(c) }, middleware.JWTWithConfig(config))
	e.PUT("/user/:id/", func(c echo.Context) error { return userHandler.UpdateUser(c) }, middleware.JWTWithConfig(config))
	e.POST("/user/signup/", func(c echo.Context) error { return userHandler.SignUp(c) })
	e.POST("/user/login", func(c echo.Context) error { return userHandler.SignIn(c) })

	// e.GET("/book/all", Bookhandler.GetBooks,middleware.JWTWithConfig(config))
	// e.GET("/book/:id", Bookhandler.GetBookByID,middleware.JWTWithConfig(config))
	// e.DELETE("/book/:id", Bookhandler.DeleteBook,middleware.JWTWithConfig(config))
	// e.POST("/book/Register", Bookhandler.Register,middleware.JWTWithConfig(config))
	// e.POST("/book/:id", Bookhandler.Reservation,middleware.JWTWithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}
