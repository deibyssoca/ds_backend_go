package actions

import (
	"net/http"
	"net/url"
	"strconv"

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
	var userID int
	if m, ok := c.Params().(url.Values); ok {
		userID, _ = strconv.Atoi(m["id"][0])
	}

	user, err := models.GetUser(userID)
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"Message": "Resource Not Found."}))
	}
	return c.Render(http.StatusOK, r.JSON(user))

}

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
