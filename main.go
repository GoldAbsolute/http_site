package main

import (
	"ex.sov/bmiddleware"
	"ex.sov/books"
	"ex.sov/db"
	"ex.sov/ex_passwords"
	exjson2 "ex.sov/exjson"
	"ex.sov/exws"
	"ex.sov/forms"
	users_pack "ex.sov/users"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//start db
	db.ConnectDatabase()
	//end db
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Welcome to my site\n")
		check(err)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		check(err)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		check(err)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		check(err)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		check(err)
	})

	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
		//get the book
		//navigate to the page
		vars := mux.Vars(request)
		title := vars["title"]
		page := vars["page"]
		_, err := fmt.Fprintf(writer, "Hello, you've requested: %s\n", title)
		check(err)
		_, err2 := fmt.Fprintf(writer, "Hello, you've requested: %s\n", page)
		check(err2)
	})
	//err := http.ListenAndServe(":8090", nil)
	//check(err)
	r.HandleFunc("/pips/{title}", books.BookHandler).Host("localhost")
	r.HandleFunc("/books/{title}", books.CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", books.ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", books.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", books.DeleteBook).Methods("DELETE")

	r.HandleFunc("/create_db", db.CreateDB).Methods("GET")

	r.HandleFunc("/secure", books.SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", books.InsecureHandler).Schemes("http")

	forms.Form1(r)
	bmiddleware.Bmiddleware(r)
	//sessions-cookie
	bmiddleware.SecretWareHandler(r)

	kickRouter := r.PathPrefix("/kick").Subrouter()
	kickRouter.HandleFunc("", AllKick)
	kickRouter.HandleFunc("/", AllKick)
	kickRouter.HandleFunc("/{id}", IdHandler)

	exjsonSubRouter := r.PathPrefix("/exjson").Subrouter()
	exjsonSubRouter.HandleFunc("/decode", exjson2.DecodeJson)
	exjsonSubRouter.HandleFunc("/encode", exjson2.EncodeJson)

	websocketSubRouter := r.PathPrefix("/ws").Subrouter()
	websocketSubRouter.HandleFunc("/", exws.WebSeeHanlder)
	websocketSubRouter.HandleFunc("/echo", exws.WebSuck)

	exPasswordSubRouter := r.PathPrefix("/pass").Subrouter()
	exPasswordSubRouter.HandleFunc("/", ex_passwords.Main)

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/read_user", users_pack.ReadUser)
	users.HandleFunc("/read_and_see_user", users_pack.ReadAndSeeUser).Methods("GET")
	users.HandleFunc("/create_user", users_pack.CreateUser)
	users.HandleFunc("/delete_user", users_pack.DeleteUser)
	users.HandleFunc("/update_user", users_pack.UpdateUser)
	users.HandleFunc("/temp_user", users_pack.TempUser)

	templates := r.PathPrefix("/templates").Subrouter()

	templates.HandleFunc("/tmpl1", tmpl_1)

	err2 := http.ListenAndServe(":8091", r)
	check(err2)

}
func tmpl_1(writer http.ResponseWriter, request *http.Request) {
	type Todo struct {
		Title string
		Done  bool
	}
	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}
	tmpl := template.Must(template.ParseFiles("template_1.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(writer, data)

}
func AllKick(writer http.ResponseWriter, request *http.Request) {

}
func IdHandler(writer http.ResponseWriter, request *http.Request) {

}

// db
