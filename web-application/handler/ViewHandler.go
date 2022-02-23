package handler

import (
	"fmt"
	"net/http"
	"reyhan/web-application/model"
)

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

	// Write respnse body
	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", page.Title, page.Body)
}

// Wiki view function handler
func GetWikiView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	title := r.URL.Path[len("/view/"):]

	page, err := model.LoadPage(title)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", page.Title, page.Body)
}
