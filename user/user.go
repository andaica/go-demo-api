package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/andaica/go-demo-api/authen"

	"github.com/gorilla/mux"
)

func RegistRouter(router *mux.Router) {
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users", getAllUsers)
	router.HandleFunc("/api/users/login", loginHandle).Methods("POST")
	router.HandleFunc("/api/user", authen.BasicAuth(getUser))
}

var UDM = UserDataMapping{}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: allUser")
	users := UDM.fetchAll()
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := getUserDataFromPost(r)
	log.Println("Endpoint hit: createUser ", user)
	if validateUser(user) == true {
		newUser, isCreateOk := UDM.insertNewUser(user)
		if isCreateOk {
			token, err := authen.CreateToken(newUser.Email, newUser.Id)
			if err == nil {
				newUser.Token = &token
				responseSuccess(w, newUser)
				return
			}
		}
	}
	responseError(w)
}

func validateUser(user User) bool {
	if len(user.Email) == 0 || len(user.Username) == 0 || len(user.Password) == 0 {
		return false
	}
	return true
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	loginData := getUserDataFromPost(r)
	log.Println("Endpoint hit: login ", loginData)
	user, isExist := UDM.getUserByLogin(loginData.Email, loginData.Password)
	if isExist {
		token, err := authen.CreateToken(user.Email, user.Id)
		if err == nil {
			user.Token = &token
			responseSuccess(w, user)
			return
		}
	}
	responseError(w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := authen.GetAuthenticatedUserId(r)
	log.Println("Endpoint hit: getUser ", userId)
	user, isExist := UDM.getUserById(userId)
	if isExist {
		responseSuccess(w, user)
	} else {
		responseError(w)
	}
}

func getUserDataFromPost(r *http.Request) User {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userReq Request
	json.Unmarshal(reqBody, &userReq)
	log.Printf("%+v\n", userReq)
	return userReq.User
}

func responseSuccess(w http.ResponseWriter, user User) {
	response := Response{convertUserResponse(user), "OK"}
	json.NewEncoder(w).Encode(response)
}

func responseError(w http.ResponseWriter) {
	response := Response{UserResponse{}, "NG"}
	json.NewEncoder(w).Encode(response)
}

func convertUserResponse(user User) UserResponse {
	var res = UserResponse{}
	v := reflect.Indirect(reflect.ValueOf(&res))
	u := reflect.Indirect(reflect.ValueOf(user))
	for i := 0; i < v.NumField(); i++ {
		f := v.FieldByName(v.Type().Field(i).Name)
		if f.CanSet() {
			f.Set(u.FieldByName(v.Type().Field(i).Name))
		}
	}
	return res
}
