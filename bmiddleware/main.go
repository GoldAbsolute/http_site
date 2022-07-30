package bmiddleware

import (
	mysessions "ex.sov/mysessions"
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

// middleware
func SecretWare(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		log.Println(request.Method)
		f(writer, request)
	}
}

// Router for handlers
func SecretWareHandler(router *mux.Router) {
	router.HandleFunc("/secret", SecretWare(mysessions.Secret))
	router.HandleFunc("/login", SecretWare(mysessions.Login))
	router.HandleFunc("/logout", SecretWare(mysessions.Logout))
}

// Route for secret
//Переместил в папку sessions
//func SecretWareRoute(writer http.ResponseWriter, request *http.Request) {
//	fmt.Fprintln(writer, "secret")
//}
