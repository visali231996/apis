package api

import (
	"database/sql"
	"net/http"
)

func Registerroute(db *sql.DB) {
	http.HandleFunc("/create", CreateHandler(db))
}
