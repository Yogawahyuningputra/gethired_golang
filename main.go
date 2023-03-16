package main

import (
	"fmt"
	"net/http"

	"backend/database"
	"backend/packages/mysql"
	"backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	
	mysql.DatabaseInit()
	database.RunMigration()
	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "3030"
	fmt.Println("Server running on port " + port)
	
	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

}
