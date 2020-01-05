package repository

import (
	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/domain/model"
)

type BookRepository interface {
	GetAll(ctx echo.Context) (book *[]model.Book, err error)
	GetByTitle(ctx echo.Context, title string) (book *model.Book, err error)
	RegisterBook(ctx echo.Context, title, content string) error
	UpdateBookInfo(ctx echo.Context, title, context string) error
	DeleteBook(ctx echo.Context, title string) error
}
