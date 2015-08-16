package views

import (
	"github.com/jfyne/accordopartners.com/back"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
)

var re *render.Render

// Get: Find the correct row in the csv
func Get(key string, data [][]string) []string {
	for i := range data {
		if data[i][0] == key {
			return data[i]
		}
	}

	return []string{}
}

func init() {
	re = render.New(render.Options{
		Layout:        "layout",
		Extensions:    []string{".html"},
		IsDevelopment: true,
		Funcs: []template.FuncMap{
			template.FuncMap{
				"get": Get,
			},
		},
	})
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
	common(w, "experience")
}

func Solutions(w http.ResponseWriter, r *http.Request) {
	common(w, "solutions")
}

func Terms(w http.ResponseWriter, r *http.Request) {
	common(w, "terms")
}
