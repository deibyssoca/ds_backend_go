package actions

import (
	"net/http"

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
	//users := []models.User{}
	users := models.GetUsers()
	return c.Render(http.StatusOK, r.JSON(users))
	//return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Users"}))

}

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

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	models.SendData(w, models.GetUsers())
// }

// GetUser a user is obtained.
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
