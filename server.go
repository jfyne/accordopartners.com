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
	router.HandleFunc("/experience", views.Experience)
	router.HandleFunc("/expos/{tag}", views.Expos)
	router.HandleFunc("/partners", views.Partners)
	router.HandleFunc("/solutions", views.Solutions)
	router.HandleFunc("/terms", views.Terms)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
