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
	app.Router.Methods(http.MethodGet).Path("/tenancies").HandlerFunc(app.getAllTenancies)
	app.Router.Methods(http.MethodGet).Path("/tenancies/{id}").HandlerFunc(app.getTenancies)
}

func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers")
	w.Header().Set("Content-Type", "application/json")

	items := make([]model.User, 0)

	cursor, err := app.Database.Query("select id, email, password, name from user")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer cursor.Close()

	for cursor.Next() {
		var item model.User
		err := cursor.Scan(&item.Id, &item.Email, &item.Password, &item.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
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

func (app *App) getTenancies(w http.ResponseWriter, r *http.Request) {
	log.Println("getTenancies")
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	items := make([]model.Tenancy, 0)

	query := "select t.id, t.user_id, t.property_id, t.start_date, t.end_date, p.address from tenancy t left join property p on p.id=t.property_id where t.user_id= " + params["id"] + " order by t.end_date asc"

	cursor, err := app.Database.Query(query)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer cursor.Close()

	for cursor.Next() {
		var item model.Tenancy
		err := cursor.Scan(&item.Id, &item.UserId, &item.PropertyId, &item.StartDate, &item.EndDate, &item.Address)
		if err != nil {
			log.Fatal(err.Error())
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

func (app *App) getAllTenancies(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllTenancies")

	w.Header().Set("Content-Type", "application/json")

	items := make([]model.Tenancy, 0)

	query := "select t.id, t.user_id, t.property_id, t.start_date, t.end_date, p.address from tenancy t left join property p on p.id=t.property_id order by t.end_date asc"

	cursor, err := app.Database.Query(query)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer cursor.Close()

	for cursor.Next() {
		var item model.Tenancy
		err := cursor.Scan(&item.Id, &item.UserId, &item.PropertyId, &item.StartDate, &item.EndDate, &item.Address)
		if err != nil {
			log.Fatal(err.Error())
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}
