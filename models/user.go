package models

import "errors"

// User struct
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Users type
type Users []User

var users = make(map[int]User)

// SetDefaultUser temp function
func SetDefaultUser() {
	user := User{ID: 1, UserName: "Deibys", Password: "1234"}
	users[user.ID] = user
}

// GetUsers get all users
func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

// GetUser by id
func GetUser(userID int) (User, error) {
	if user, ok := users[userID]; ok {
		return user, nil
	}
	return User{}, errors.New("The user is not on the map")
}
