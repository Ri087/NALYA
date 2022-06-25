package main

import (
	"nalya/nalya"
	"net/http"
)

// type App struct {
// 	Router *mux.Router
// 	DB     *sql.DB
// }

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/getData", nalya.DataHandler)
	http.ListenAndServe(":8080", nil)
}
