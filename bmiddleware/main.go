package bmiddleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func foo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "foo")
}

func bar(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "bar")
}

func Bmiddleware(r *mux.Router) {
	r.HandleFunc("/foo", logging(foo))
	r.HandleFunc("/bar", logging(bar))
}
