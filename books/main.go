package books

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {

}
func ReadBook(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello, you've requested URL: %s\n", request.URL.Path)
	check(err)
	vars := mux.Vars(request)
	title := vars["title"]
	_, err2 := fmt.Fprintf(writer, "Hello, you've requested title: %s\n", title)
	check(err2)
}
func UpdateBook(writer http.ResponseWriter, request *http.Request) {

}
func DeleteBook(writer http.ResponseWriter, request *http.Request) {

}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "hui")
	check(err)
}
func SecureHandler(w http.ResponseWriter, r *http.Request) {

}
func InsecureHandler(w http.ResponseWriter, r *http.Request) {

}
