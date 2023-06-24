package main

import (
	"log"
	"net/http"
	c "web/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.Page)
	mux.HandleFunc("/ascii-art", c.Ascii)
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))
	log.Println("Запуск веб-сервера на http://localhost:8080/")
	err := http.ListenAndServe(":8080", mux)
	log.Println("You can`t ", err)
}
