package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST        = "localhost"
	CONN_PORT        = "8080"
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:password@/go_db"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database ::", connectionError)
	}
}

func getCurrentDb(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATABASE() as  db")
	if err != nil {
		log.Print("error executing query:: ", err)
		return
	}
	var db string
	for rows.Next() {
		rows.Scan(&db)
	}
	fmt.Fprintf(w, "Current database is :: %s", db)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getCurrentDb).Methods("GET")
	defer db.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}
