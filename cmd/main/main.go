package main

import (
	"log"
	"net/http"

	"github.com/bhagyadeep2002/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//create a router and pass it to routes file
//start a http server at localhost:8000
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
