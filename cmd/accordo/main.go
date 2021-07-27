package main

import (
	"github.com/gorilla/mux"
	"github.com/jfyne/accordopartners.com/views"
	"github.com/urfave/negroni"
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
	router.HandleFunc("/reports", views.Reports).Methods("GET")
	router.HandleFunc("/reports", views.ReportsSend).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
