package controllers

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {
	response := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: error,
	}
	w.WriteHeader(code)
	tmpl, _ := template.ParseFiles("ui/html/error.html")
	tmpl.Execute(w, response)
}
