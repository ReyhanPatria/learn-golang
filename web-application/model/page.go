package model

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

var pathPrefix = "page/"

func (p *Page) Save() error {
	filename := pathPrefix + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := pathPrefix + title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
