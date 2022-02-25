package handler

import (
	"html/template"
	"net/http"
	"reyhan/web-application/model"
)

var templatePrefix string = "./page/templates/"

func renderTemplate(w http.ResponseWriter, templateName string, page *model.Page) error {
	// Create html template for page
	templateFilename := templatePrefix + templateName + ".html"
	t, err := template.ParseFiles(templateFilename)

	// Handle if template is not found
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	// Apply Page to html template
	return t.Execute(w, page)
}

func isRequestMethodValid(w http.ResponseWriter, actual string, expected ...string) bool {
	for _, method := range expected {
		if method == actual {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return true
		}
	}
	return false
}

func loadPage(w http.ResponseWriter, title string) (*model.Page, error) {
	page, err := model.LoadPage(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return nil, err
	}
	return page, nil
}

// Index function handler
func GetIndexView(w http.ResponseWriter, r *http.Request) {
	// Handle if http method is not GET
	if !isRequestMethodValid(w, r.Method, http.MethodGet) {
		return
	}

	// Get page file
	if page, err := loadPage(w, "index"); err == nil {
		renderTemplate(w, "blog", page)
	}
}

// Blog view function handler
func GetBlogView(w http.ResponseWriter, r *http.Request) {
	// Check if http method is GET
	if !isRequestMethodValid(w, r.Method, http.MethodGet) {
		return
	}

	// Get page title
	title := r.URL.Path[len("/view/"):]

	// Get Page
	if page, err := loadPage(w, title); err == nil {
		renderTemplate(w, "blog", page)
	}
}
