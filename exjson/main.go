package exjson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func DecodeJson(writer http.ResponseWriter, request *http.Request) {
	user := User{
		Firstname: "first_name_decode",
		Lastname:  "last-name_decode",
		Age:       100,
	}
	json.NewDecoder(request.Body).Decode(&user)
	//json.NewEncoder(writer).Encode(user)

	fmt.Fprintf(writer, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}
func EncodeJson(writer http.ResponseWriter, request *http.Request) {
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	json.NewEncoder(writer).Encode(peter)
}
