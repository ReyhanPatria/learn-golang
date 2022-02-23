package handler

import (
	"html/template"
	"net/http"
	"reyhan/web-application/model"
)

var templatePrefix string = "./page/templates/"

// Index function handler
func GetIndexView(w http.ResponseWriter, r *http.Request) {
	// Handle if http method is not GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get page file
	page, err := model.LoadPage("index")

	// Handle if page file doesn't exist
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Create html template for page
	templateFilename := templatePrefix + "wiki.html"
	t, err := template.ParseFiles(templateFilename)

	// Handle if template is not found
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Apply Page to html template
	t.Execute(w, page)
}

// Wiki view function handler
func GetWikiView(w http.ResponseWriter, r *http.Request) {
	// Check if http method is GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get page title
	title := r.URL.Path[len("/view/"):]

	// Get Page
	page, err := model.LoadPage(title)

	// Handle if page file doesn't exist
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Create html template for page
	templateFilename := templatePrefix + "wiki.html"
	t, err := template.ParseFiles(templateFilename)

	// Handle if template is not found
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Apply Page to html template
	t.Execute(w, page)
}
