package views

import (
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
)

type report struct {
	Title string
	URL   string
}

type reportDownload struct {
	Email  string
	Report string `datastore:",noindex"`
}

var reports map[string]report

func init() {
	reports = map[string]report{
		"iyq-2021": {
			Title: "IYQ Project Implementation Scenario - Investment Perspective Summer 2021",
			URL:   "https://storage.googleapis.com/accordo-partners-reports/IYQ%20Project%20Implementation%20Scenario%20-%20Investment%20Perspective%20Summer%202021.pdf",
		},
	}
}

func getReport(key string) report {
	r, ok := reports[key]
	if !ok {
		return report{}
	}
	return r
}

func Reports(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{
		"Error":    "",
		"Download": "",
	}
	re.HTML(w, http.StatusOK, "reports", context)
}

func ReportsSend(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{
		"Error":    "",
		"Download": false,
	}

	rd := &reportDownload{
		Email:  r.FormValue("email"),
		Report: r.FormValue("report"),
	}

	key := datastore.NameKey("reportDownload", rd.Email, nil)
	report := getReport(rd.Report)
	context["Download"] = true
	context["Report"] = report

	if store == nil {
		re.HTML(w, http.StatusOK, "reports", context)
		return
	}
	if _, err := store.Put(r.Context(), key, rd); err != nil {
		log.Println(err)
		context["Error"] = "Something went wrong, please try again"
		re.HTML(w, http.StatusInternalServerError, "reports", context)
		return
	}

	re.HTML(w, http.StatusOK, "reports", context)
}
