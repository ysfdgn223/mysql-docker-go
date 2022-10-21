package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	GetDB() *sql.DB
}

type db struct {
	d *sql.DB
}

// After the connection is made, don`t forget to close db -> defer db.Close()`
func NewDBConnection(dbc Config) DB {
	dbmsName, str := dbc.GetDBConnectionString()
	currentDB, err := sql.Open(dbmsName, str)
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	return &db{d: currentDB}
}

func (d *db) GetDB() *sql.DB {
	return d.d
}
