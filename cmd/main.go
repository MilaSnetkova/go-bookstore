package main

import (
	"log"
	"net/http"

	"github.com/MilaSnetkova/go-bookstore/pkg/routers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	r := mux.NewRouter()
	routers.BookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

	
}