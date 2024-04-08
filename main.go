package main

import (
	"database/sql"
	"dockerdcc/database"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome!!!</h1>")
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome 2</h1>")
	var err error

	query := "SELECT * FROM MsApakek WHERE ID = 1"
	result, err := db.Query(query)
	if !result.Next() {
		fmt.Fprint(w, "No data with productId 001")
		return
	}
	if err != nil {
		log.Fatal("Error Query!")
	}
	defer result.Close()
	apakek := database.MsApakek{}
	err = result.Scan(&apakek.ID, &apakek.Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%d %s", apakek.ID, apakek.Name)
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:sasasa12@tcp(localhost:3306)/apakek")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/abc", handlerFunc2)
	fmt.Println("starting the server...")
	http.ListenAndServe(":8080", nil)
}
