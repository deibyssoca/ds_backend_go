package actions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deibyssoca/ds_backend_go/models"
	"github.com/gobuffalo/buffalo"
)

// "net/url"
// "strconv"

// var db = make(map[uuid.UUID]models.User)

// UserResource handler
type UserResource struct{}

// GetUsers all users are obtained.
func (ur UserResource) GetUsers(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(models.GetUsers()))
}

// GetUser by id
func (ur UserResource) GetUser(c buffalo.Context) error {

	if m, ok := c.Params().(url.Values); ok {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Get User!"}))
}

// // GetUser a user is obtained.
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	// Add the header
// 	w.Header().Set("Content-type", "application/json")
// 	//user := models.User{ID: 1, UserName: "Deibys", Password: "1234"}
// 	vars := mux.Vars(r)
// 	userID, _ := strconv.Atoi(vars["id"])

// 	if user, err := models.GetUser(userID); err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		models.SendNotFound(w)
// 	} else {
// 		models.SendData(w, user)
// 	}
// }

// // CreateUser user created
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "User created.!")
// }

// // UpdateUser user updated.
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "User updated!")
// }

// // DeleteUser user deleted
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "User deleted!")
// }

// func (ur UserResource) List(c buffalo.Context) error {
// 	tx, err := pop.Connect("development")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	users := []models.User{}
// 	error := tx.All(&users)
// 	if err != nil {
// 		log.Panic(error)
// 	}
// 	return c.Render(200, r.JSON(users))
// }
