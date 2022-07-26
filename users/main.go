package users

import (
	"ex.sov/db"
	"fmt"
	"net/http"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadUser(writer http.ResponseWriter, request *http.Request) {
	db := db.ReturnDB()
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	//SQL text SELECT id, username, password, created_at FROM users WHERE id = ?
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	row := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)

	fmt.Fprintf(writer, "Row: ", row)
}
func ReadAndSeeUser(writer http.ResponseWriter, request *http.Request) {
	_, err6 := fmt.Fprint(writer, "Op\n\n")
	check(err6)
	db := db.ReturnDB()
	//result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	//check(err)
	//liid64, err := result.LastInsertId()
	//check(err)
	//ra64, err2 := result.RowsAffected()
	//check(err2)
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	//needId := 3
	//SQL text SELECT id, username, password, created_at FROM users WHERE id = ?
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	check(err)
	_, err5 := fmt.Fprint(writer, "Первая строка юзеров: ", "\n", id, "\n", username, "\n", password, "\n", createdAt, "\n")
	check(err5)
	// Trying...
	_, err7 := fmt.Fprint(writer, "Все строки юзеров: ", "\n")
	check(err7)
	// Все строки юзеров
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	defer rows.Close()
	check(err)
	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		check(err)
		users = append(users, u)
	}
	err8 := rows.Err()
	check(err8)
	// Вывод
	for i, u := range users {
		fmt.Fprintf(writer, "%v: UserId: %d, UserName: %s, UserPassword: %s, CreatedDate: %v \n", i, u.id, u.username, u.password, u.createdAt)
	}
}
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	db := db.ReturnDB()
	username := "johndoe6"
	password := "secret6"
	createdAt := time.Now()
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	check(err)
	liid64, err := result.LastInsertId()
	check(err)
	ra64, err2 := result.RowsAffected()
	check(err2)
	_, err3 := fmt.Fprintf(writer, "Последнее добавленое Id(сейчас добавилось): %d \nЧисло затронутых строк: %d \nПользователь %s успешно создан!\n", liid64, ra64, username)
	check(err3)

}
func DeleteUser(writer http.ResponseWriter, request *http.Request) {

}
func UpdataUser(writer http.ResponseWriter, request *http.Request) {

}
