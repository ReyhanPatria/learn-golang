package handler

import (
	"fmt"
	"net/http"
	"reyhan/web-application/model"
)

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
