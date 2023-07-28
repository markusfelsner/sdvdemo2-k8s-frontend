package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost     = "mysql-service"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = "Sommer123456"
	dbName     = "sdvdemodatabase"
)

func main() {
	// Verbinde mit der MySQL-Datenbank
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Erstelle den HTTP-Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM sdvdemotabelle")
		if err != nil {
			fmt.Fprintf(w, "SDVDemo mit Image markusfelsner/frontend:1.0")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var id int
		var name string
		var age int

		var result string

		for rows.Next() {
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result += fmt.Sprintf("ID: %d, Name: %s, Age: %d\n", id, name, age)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, result)
	})

	// Starte den HTTP-Server
	log.Println("Server gestartet, höre auf Port 8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
