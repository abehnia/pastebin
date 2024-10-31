package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	psqlInfo := os.Getenv("DB_STRING")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO metric (timestamp, name, fields, tags)
		VALUES ($1, $2, $3, $4)
		`

	_, err := db.Exec(sqlStatement, metricRow.Timestamp, metricRow.Name, fields, tags)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/hello", HelloWorld)
	log.Fatal(http.ListenAndServe(":8000", nil))
}