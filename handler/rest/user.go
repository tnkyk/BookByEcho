package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/tnkyk/BookByEcho/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Index(ctx echo.Context) (err error)
	SignUp(ctx echo.Context) (err error)
	SignIn(ctx echo.Context) (err error)
	UpdateUser(ctx echo.Context) (err error)
	DeleteUser(ctx echo.Context) (err error)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (uh *userHandler) Index(ctx echo.Context) (err error) {
	type UserField struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Favorite string `json:"favorite"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Users []UserField `json:"users"`
	}

	users, err := uh.userUsecase.UserGetAll(ctx)
	if err != nil {
		log.Println(err)
	}
	res := new(response)
	uf := UserField{}
	for _, user := range *users {
		uf.Id = string(user.ID)
		uf.Name = user.Name
		uf.Email = user.Email
		uf.Favorite = user.Favorite
		res.Users = append(res.Users, uf)
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(ctx.Response()).Encode(res); err != nil {
		http.Error(ctx.Response(), "Internal Server Error", 500)
		return
	}

}

func (uh *userHandler) SignUp(ctx echo.Context) (err error) {
	type UserField struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	id, err := uuid.NewRandom()
	pwd := ctx.Request().FormValue("password")
	hashpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	name := ctx.Request().FormValue("name")
	email := ctx.Request().FormValue("email")
	var createdAt time.Time = time.Now()
	user := &UserField{
		Id:    id.String(),
		Name:  name,
		Email: email,
	}
	err = uh.userUsecase.SignUp(ctx, id.String(), name, email, string(hashpwd), "", &createdAt, &createdAt)
	if err != nil {
		log.Println(err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		ctx.Response().Write([]byte("Don't create your user"))
		return
	}
	ans, err := json.Marshal(user)
	ctx.Response().Write(ans)
	ctx.Response().WriteHeader(http.StatusOK)
	ctx.Response().Header().Set("Content-Type", "apprication/json")
}

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create the Signin handler
func (uh *userHandler) SignIn(ctx echo.Context) (err error) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(ctx.Request().Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	user, err := uh.userUsecase.GetByName(ctx, creds.Username)

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if err != nil || user.Password != creds.Password {
		ctx.Response().WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			Id:        user.ID,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(ctx.Response(), &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func (uh *userHandler) UpdateUser(ctx echo.Context) (err error) {
	id := ctx.Request().FormValue("id")
	name := ctx.Request().FormValue("name")
	password := ctx.Request().FormValue("passowrd")
	email := ctx.Request().FormValue("email")
	favorite := ctx.Request().FormValue("favorite")
	var updatedAt time.Time = time.Now()

	user, err := uh.userUsecase.UserUpdate(ctx, id, name, email, password, favorite, &updatedAt)
	if err != nil {
		log.Println(err)
		return
	}
	if err = json.NewEncoder(ctx.Response()).Encode(user); err != nil {
		ctx.Response().Status = 500
		return
	}

}

func (uh *userHandler) DeleteUser(ctx echo.Context) (err error) {
	name := ctx.Request().FormValue("name")
	err := uh.userUsecase.DeleteUser(ctx, name)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.Response().Write([]byte("Delete user" + name))
	ctx.Response().WriteHeader(http.StatusOK)
	ctx.Response().Header().Set("Content-Type", "apprication/json")
}
