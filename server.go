package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jfyne/accordopartners.com/views"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", views.Home)
	router.HandleFunc("/about", views.About)
	router.HandleFunc("/contact", views.Contact)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
