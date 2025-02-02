module github.com/GrantCanty/testing

go 1.23.2

require (
	github.com/gorilla/mux v1.8.1
	github.com/rs/cors v1.11.1
)

replace github.com/GrantCanty/testing/appContext => ../appContext
replace github.com/GrantCanty/testing/routes => ../routes
replace github.com/GrantCanty/testing/types => ../types