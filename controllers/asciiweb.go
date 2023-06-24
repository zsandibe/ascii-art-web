package controllers

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	f "web/internal"
)

type Export struct {
	Result string
	Text   string
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
	}
	if !(r.Form.Has("format") && r.Form.Has("text")) {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	Text := r.FormValue("text")
	if len(Text) > 400 {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	format := r.FormValue("format")
	arrByte, err := ioutil.ReadFile("./fonts/" + format + ".txt")
	if err != nil {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	hash := f.HashMD5(string(arrByte))
	if !(f.CheckHash(hash)) {
		Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	mapChar := f.MakeMap(arrByte)
	word, err := f.SplitWords(Text)
	if err != nil {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	Result := f.Print(word, mapChar)
	Output := Export{Result: Result, Text: Text}
	// export := r.FormValue("download")
	if r.FormValue("process") == "download" {
		a := strings.NewReader(Result)
		w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
		w.Header().Set("Content-Length", strconv.Itoa(len(Result)))
		io.Copy(w, a)
	}
	tmpl, err := template.ParseFiles("ui/html/template.html")
	tmpl.Execute(w, Output)
}
