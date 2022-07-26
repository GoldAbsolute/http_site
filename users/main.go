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
	//SQL text
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
