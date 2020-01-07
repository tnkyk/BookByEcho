package usecase

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/domain/model"
	"github.com/tnkyk/BookByEcho/domain/repository"
)

type UserUsecase interface {
	UserGetAll(ctx echo.Context) (*[]model.User, error)
	GetByName(ctx echo.Context, name string) (*model.User, error)
	UserUpdate(ctx echo.Context, id string, name string, email string, password string, favorite string, updatedAt *time.Time) (*model.User, error)
	SignUp(ctx echo.Context, id string, name string, email string, password string, favorite string, createdAt, updatedAt *time.Time) error
	DeleteUser(ctx echo.Context, name string) error
}
type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) UserGetAll(ctx echo.Context) (*[]model.User, error) {
	users, err := uu.userRepository.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) GetByName(ctx echo.Context, name string) (*model.User, error) {
	user, err := uu.userRepository.GetByName(ctx, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, err
}

func (uu *userUsecase) UserUpdate(ctx echo.Context, id string, name string, email string, password string, favorite string, updatedAt *time.Time) (*model.User, error) {
	user, err := uu.userRepository.Update(ctx, id, name, email, password, favorite, updatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (uu *userUsecase) SignUp(ctx echo.Context, id, name, email, password, favorite string, createdAt, updatedAt *time.Time) error {
	err := uu.userRepository.SignUp(ctx, id, name, email, password, favorite, createdAt, updatedAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (uu *userUsecase) DeleteUser(ctx echo.Context, name string) error {
	err := uu.userRepository.Delete(ctx, name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
