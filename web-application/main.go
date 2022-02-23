package main

import (
	"net/http"
	"reyhan/web-application/handler"
)

func main() {
	http.HandleFunc("/", handler.GetIndexView)
	http.HandleFunc("/view/", handler.GetWikiView)

	http.ListenAndServe("localhost:8080", nil)
}
