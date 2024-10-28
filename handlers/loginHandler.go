package handlers

import (
	"crud-repo-2/f"
	"crud-repo-2/middleware"
	"crud-repo-2/requests"
	"crud-repo-2/responses"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

// -------------- JWT ----------------
// var validate *validator.Validate

const USERNAME = middleware.USERNAME
const PASSWORD = middleware.PASSWORD

type tokenJwtHandler struct {
	token responses.TokenJWTResponse
	user  requests.User
}

func NewTokenJwtHandler() *tokenJwtHandler {
	return &tokenJwtHandler{}
}

func (j *tokenJwtHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var validate *validator.Validate
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	// var userRequest requests.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&j.user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(j.user)

	if err != nil {
		f.WriteToJsonError(w, r, f.ErrorValidation(err))
		return
	}

	if j.user.Username != USERNAME || j.user.Password != PASSWORD {
		f.WriteToJsonError(w, r, "Username or Password is incorrect")
		return
	}

	tokenString, err := j.CreateToken()
	if err != nil {
		f.WriteToJsonError(w, r, err.Error())
		return
	}

	token := responses.TokenJWTResponse{
		Token: tokenString,
	}

	j.token = token
	f.WriteToJson(w, r, token)

}

func (j *tokenJwtHandler) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": j.user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
