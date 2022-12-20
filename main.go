package main

import (
	"BE-DeweTour/database"
	"BE-DeweTour/pkg/mysql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init godotenv

	//initial DB
	mysql.DatabaseInit()

	//run migration
	database.RunMigration()

	//initial routes
	r := mux.NewRouter()

	//initial group version route

	//inititalization static route path

	fmt.Println("server running on port:5000")
	http.ListenAndServe(":5000", r)
}
