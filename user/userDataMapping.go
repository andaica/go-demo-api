package user

import (
	"fmt"
	"strconv"

	"go-demo-api/db"

	_ "github.com/go-sql-driver/mysql"
)

const table = "user"

type UserDataMapping struct{}

func (u UserDataMapping) fetchAll() []User {
	users := []User{}
	query := "select * from user"
	result, _ := db.Execute(query)

	for result.Next() {
		user := User{}

		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Image, &user.Token)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func (u UserDataMapping) insertNewUser(user User) (newUser User, isOk bool) {
	var image string = ""
	if user.Image != nil {
		image = "'" + *user.Image + "'"
	} else {
		image = "null"
	}
	query := fmt.Sprintf(
		"insert into user ( username, email, password, bio, image ) values ( '%s', '%s', '%s', '%s', %s )",
		user.Username, user.Email, user.Password, user.Bio, image,
	)
	_, err := db.Execute(query)

	if err != nil {
		return newUser, false
	} else {
		newUser = getNewUser(user)
	}
	return newUser, true
}

func getNewUser(user User) (newUser User) {
	query := "select * from user where email = '" + user.Email + "' and password = '" + user.Password + "'"
	result, _ := db.Execute(query)

	if result.Next() {
		err := result.Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Bio, &newUser.Image, &newUser.Token)
		if err != nil {
			panic(err.Error())
		}
	}

	return newUser
}

func (u UserDataMapping) getUserByLogin(email, password string) (user User, isExist bool) {
	query := "select * from user where email = '" + email + "' and password = '" + password + "'"
	result, _ := db.Execute(query)

	if result.Next() {
		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Image, &user.Token)
		if err != nil {
			panic(err.Error())
		}
		isExist = true
	} else {
		isExist = false
	}

	return user, isExist
}

func (u UserDataMapping) getUserById(id uint) (user User, isExist bool) {
	query := "select * from user where id = " + strconv.Itoa(int(id))
	result, _ := db.Execute(query)

	if result.Next() {
		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Image, &user.Token)
		if err != nil {
			panic(err.Error())
		}
		isExist = true
	} else {
		isExist = false
	}

	return user, isExist
}
