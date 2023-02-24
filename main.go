package main

import (
	"fmt"
	"net/http"

	"github.com/JANCARLO123/go-gorm/routes"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Hello word")

	r:= mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":5000",r)
}
