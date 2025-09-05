package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	dsn := "username:password@tcp(localhost)dbname?parseTime=True" //db file info
	db, err := sql.Open("msql", dsn)                               //prepares the drivers and validate if dsn parameters are correct
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.Registerroute(db)

	//http server
	log.Println("server starting on port 9090.......")
	log.Fatal(http.ListenAndServe(".9090", nil))

}

//command to check which ports are listening windows: netstat -aon
