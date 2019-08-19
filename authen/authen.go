package authen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

const BEARER_SCHEMA = "Bearer "

func CreateToken(email string, id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = fmt.Sprint(id)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

type Handler func(w http.ResponseWriter, r *http.Request)

func BasicAuth(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		token := getTokenFromHeader(r)
		if len(token) == 0 {
			response := map[string]interface{}{"status": "NG", "message": "Invalid Authorization"}
			json.NewEncoder(w).Encode(response)
			return
		} else {
			claims, err := verifyToken(token)
			// log.Println("data-auth: ", claims)
			if err != nil {
				response := map[string]interface{}{"status": "NG", "message": "Error verifying JWT token"}
				json.NewEncoder(w).Encode(response)
				return
			}

			id := claims.(jwt.MapClaims)["id"].(string)
			r.Header.Set("user_id", id)
			handler(w, r)
		}
	}
}

func getTokenFromHeader(r *http.Request) (token string) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) == 0 {
		return ""
	}
	token = authorization[len(BEARER_SCHEMA):]
	return token
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func GetAuthenticatedUserId(r *http.Request) uint {
	user_id := r.Header.Get("user_id")
	u64, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		panic(err)
	}
	id := uint(u64)
	return id
}
