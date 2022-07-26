module ex.sov/users

go 1.19

replace ex.sov/db => ../db

require ex.sov/db v0.0.0-00010101000000-000000000000

require github.com/go-sql-driver/mysql v1.6.0 // indirect
