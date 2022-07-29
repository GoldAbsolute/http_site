package forms

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func Form1(r *mux.Router) {
	tmpl := template.Must(template.ParseFiles("./forms/form.html"))
	r.HandleFunc("/form1", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
			return
		}

		details := ContactDetails{
			Email:   request.FormValue("email"),
			Subject: request.FormValue("subject"),
			Message: request.FormValue("message"),
		}
		_ = details

		tmpl.Execute(writer, struct {
			Success bool
		}{true})
	})
}
