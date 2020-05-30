package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darren-rose/golang-mysql/app"
	"github.com/darren-rose/golang-mysql/db"
	"github.com/darren-rose/golang-mysql/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.Println("Starting")

	var c model.Configuration
	err := envconfig.Process("GOLANG_MYSQL", &c)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	database, err := db.CreateDatabase(c)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
		return
	}

	defer database.Close()

	log.Println("Connected to database")

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Println(fmt.Sprintf("Listening on port %d", c.AppPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.AppPort), app.Router))

}
