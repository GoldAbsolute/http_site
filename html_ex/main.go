package html_ex

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//tmpl, err := template.ParseFiles("layout.html")
	// or
	//tmpl := template.Must(template.ParseFiles("layout.html"))
	//check(err)
	//var HandlerTemplate = func(w http.ResponseWriter, r *http.Request) {
	//	tmpl.Execute(w, "data goes here")
	//}
}
