module ex.sov/bmiddleware

go 1.19

require (
	ex.sov/mysessions v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)

replace ex.sov/mysessions => ../mysessions
