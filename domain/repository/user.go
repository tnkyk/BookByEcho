package repository

import (
	"time"

	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/domain/model"
)

type UserRepository interface {
	GetAll(ctx echo.Context) (*[]model.User, error)
	GetByName(ctx echo.Context, name string) (*model.User, error)
	SignUp(ctx echo.Context, id, name, email, password, favorite string, createdAt, updatedAt *time.Time) error
	Update(ctx echo.Context, id, name, email, password, favorite string, updatedAt *time.Time) (*model.User, error)
	Delete(ctx echo.Context, name string) error
}
