package main

import (
	"net/http"
	"text/template"
)

func main() {

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, "oops \n", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	mux.HandleFunc("GET /whygotth", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/whygotth.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	http.ListenAndServe(":8080", mux)
}
