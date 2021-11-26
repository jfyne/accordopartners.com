package views

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"github.com/jfyne/accordopartners.com/back"
	"github.com/unrolled/render"
)

var re *render.Render
var store *datastore.Client

// Get: Find the correct row in the csv
func Get(key string, data [][]string) []string {
	for i := range data {
		if len(data[i]) == 0 {
			continue
		}
		if data[i][0] == key {
			return data[i]
		}
	}

	return []string{}
}

func init() {
	re = render.New(render.Options{
		Directory:     "front/templates",
		Layout:        "layout",
		Extensions:    []string{".html"},
		IsDevelopment: true,
		Funcs: []template.FuncMap{
			{
				"get": Get,
			},
		},
	})

	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		return
	}

	client, err := datastore.NewClient(context.Background(), "compliance-167400")
	if err != nil {
		log.Fatal(err)
	}
	store = client
}

func common(w http.ResponseWriter, tmpl string) {
	data, _ := back.SiteCopy()
	context := map[string]interface{}{
		"Data": data,
	}
	re.HTML(w, http.StatusOK, tmpl, context)
}

func Home(w http.ResponseWriter, r *http.Request) {
	common(w, "index")
}

func About(w http.ResponseWriter, r *http.Request) {
	common(w, "about")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	common(w, "contact")
}

func Experience(w http.ResponseWriter, r *http.Request) {
	data, _ := back.SiteCopy()
	categories, _ := back.AllCategories()
	context := map[string]interface{}{
		"Data":       data,
		"Categories": categories,
	}
	re.HTML(w, http.StatusOK, "experience", context)
}

func Solutions(w http.ResponseWriter, r *http.Request) {
	common(w, "solutions")
}

func Terms(w http.ResponseWriter, r *http.Request) {
	common(w, "terms")
}

func Partners(w http.ResponseWriter, r *http.Request) {
	partners, _ := back.Partners()
	context := map[string]interface{}{
		"Partners": partners,
	}
	re.HTML(w, http.StatusOK, "partners", context)
}

func Expos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]

	categories, _ := back.AllCategories()
	expos, _ := back.Expos(tag)

	context := map[string]interface{}{
		"Category":   tag,
		"Categories": categories,
		"Expos":      expos,
	}
	re.HTML(w, http.StatusOK, "expos", context)
}
