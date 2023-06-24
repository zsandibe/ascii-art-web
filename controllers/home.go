package controllers

import (
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("ui/html/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
