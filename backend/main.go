package main

import (
	"database/sql"
	"fmt"
	"log"
	"nalya/nalya"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// type App struct {
// 	Router *mux.Router
// 	DB     *sql.DB
// }
func initDataBase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	SqltStmt := `
	CREATE TABLE IF NOT EXISTS users (
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					name TEXT NOT NULL,
					email TEXT NOT NULL,
					password TEXT NOT NULL
				);
				`
	_, err = db.Exec(SqltStmt)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func insertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO users (name,email,password) VALUES (?,?,?)`, name, email, password)
	return result.LastInsertId()

}

func main() {

	db := initDataBase("NALYA_DB.db")
	fmt.Println(db)

	insertIntoUsers(db, "admin", "admin@admin.admin", "azeaze")

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/getData", nalya.DataHandler)
	http.ListenAndServe(":8080", nil)
}
