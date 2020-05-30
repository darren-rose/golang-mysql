package app

import (
	"database/sql"
	"encoding/json"
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
	app.Router.
		Methods("GET").
		Path("/users").
		HandlerFunc(app.getUsers)
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
