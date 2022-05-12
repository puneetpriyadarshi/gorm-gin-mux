package main

import (
	"log"
	"net/http"
	"root/routes"
	"root/services"
	"root/utility"
)

func main() {

	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
