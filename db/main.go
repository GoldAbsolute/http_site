package db

import (
	"database/sql"
	"fmt"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/http_db?parseTime=true")
	check(err)
	err2 := db.Ping()
	check(err2)
}
func ReturnDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/http_db?parseTime=true")
	check(err)
	err2 := db.Ping()
	check(err2)
	return db
}
func CreateDB(writer http.ResponseWriter, request *http.Request) {
	db := ReturnDB()
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err := db.Exec(query)
	check(err)

	_, err3 := fmt.Fprint(writer, "База данных успешно создана!")
	check(err3)
}
