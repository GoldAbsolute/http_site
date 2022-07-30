module example.com

go 1.19

require (
	ex.sov/bmiddleware v0.0.0-00010101000000-000000000000
	ex.sov/books v0.0.0-00010101000000-000000000000
	ex.sov/db v0.0.0-00010101000000-000000000000
	ex.sov/forms v0.0.0-00010101000000-000000000000
	ex.sov/users v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	ex.sov/mysessions v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)

replace ex.sov/books => ./books

replace ex.sov/db => ./db

replace ex.sov/users => ./users

replace ex.sov/html-ex => ./html-ex

replace ex.sov/forms => ./forms

replace ex.sov/bmiddleware => ./bmiddleware

replace ex.sov/sessions => ./mysessions

replace ex.sov/mysessions => ./mysessions
