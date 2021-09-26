package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/ngamux/ngamux"
)

//go:embed  templates/index.html
var IndexTemplate embed.FS

//go:embed  templates/blog.html
var BlogTemplate embed.FS

func IndexController(rw http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFS(IndexTemplate, "templates/index.html")
	if err != nil {
		return ngamux.StringWithStatus(rw, http.StatusInternalServerError, err.Error())
	}
	tmpl.Execute(rw, []string{})
	return nil
}

func BlogController(rw http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFS(BlogTemplate, "templates/blog.html")
	if err != nil {
		return ngamux.StringWithStatus(rw, http.StatusInternalServerError, err.Error())
	}

	data, err := GetAllUsersPosts()
	if err != nil {
		return ngamux.StringWithStatus(rw, http.StatusInternalServerError, err.Error())
	}

	tmpl.Execute(rw, struct {
		Posts []Post
	}{data})
	return nil
}
