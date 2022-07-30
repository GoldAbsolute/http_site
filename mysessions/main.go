package sessions

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Secret(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "cookie-name")
	if err != nil {
		panic(err)
	}
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(writer, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(writer, "Cake is a lie!")
}
func Login(writer http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "cookie-name")
	if err != nil {
		panic(err)
	}
	session.Values["authenticated"] = true
	session.Save(request, writer)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		panic(err)
	}
	session.Values["authenticated"] = false
	session.Save(r, w)
}
