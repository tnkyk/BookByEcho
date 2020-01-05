package persistence

import (
	"time"

	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/config"
	"github.com/tnkyk/BookByEcho/domain/model"
	"github.com/tnkyk/BookByEcho/domain/repository"
)

type UserPersistence struct {
}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}

func (up *UserPersistence) GetAll(ctx echo.Context) (*[]model.User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		ctx.Logger().Error(err)
	}
	users := []model.User{}
	user := model.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Favorite)
		if err != nil {
			ctx.Logger().Error(err)
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (up *UserPersistence) GetByName(ctx echo.Context, name string) (*model.User, error) {
	row := config.DB.QueryRow("SELECT * FROM users WHERE name = ?", name)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Favorite)
	if err != nil {
		ctx.Logger().Error(err)
		return nil, err
	}
	return &user, nil
}

func (up *UserPersistence) SignUp(ctx echo.Context, id, name, email, password, favorite string, createdAt, updatedAt *time.Time) error {
	stmt, err := config.DB.Prepare(`INSERT INTO users (id, name,email,password,favorite) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}
	_, err = stmt.Exec(id, name, email, password, favorite)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}
	return nil
}

func (up *UserPersistence) Update(ctx echo.Context, id, name, email, password, favorite string, updatedAt *time.Time) (*model.User, error) {
	stmt, err := config.DB.Prepare("UPDATE users SET name = ?,email= ?,password=?,favorite=? WHERE id = ?")
	if err != nil {
		ctx.Logger().Errorf("can't update by sql Prepare %v", err)
		return nil, err
	}
	_, err = stmt.Exec(name, email, password, favorite, id)
	if err != nil {
		ctx.Logger().Errorf("can't update by sql Exec %v", err)
		return nil, err
	}
	user := model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		Favorite:  favorite,
		UpdatedAt: updatedAt,
	}
	return &user, nil
}

func (up *UserPersistence) Delete(ctx echo.Context, id string) error {
	stmt, err := config.DB.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		ctx.Logger().Errorf("can't delete by sql Prepare %v", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		ctx.Logger().Errorf("can't Exec by sql Exec %v", err)
		return err
	}
	return nil
}
