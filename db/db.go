package db

import (
	"database/sql"
	"fmt"

	"github.com/darren-rose/golang-mysql/model"
)

func CreateDatabase(c model.Configuration) (*sql.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
