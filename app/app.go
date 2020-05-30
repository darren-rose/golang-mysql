package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darren-rose/golang-mysql/model"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.Methods(http.MethodGet).Path("/users").HandlerFunc(app.getUsers)
	app.Router.Methods(http.MethodDelete).Path("/users/{id}").HandlerFunc(app.deleteUser)
}

func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers")
	w.Header().Set("Content-Type", "application/json")

	var users []model.User

	cursor, err := app.Database.Query("select id, email, password, name from user")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer cursor.Close()

	for cursor.Next() {
		var user model.User
		err := cursor.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteUser")
	params := mux.Vars(r)
	stmt, err := app.Database.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	rows, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	log.Println(fmt.Sprintf("deleted %d rows", rows))
	if rows == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}
