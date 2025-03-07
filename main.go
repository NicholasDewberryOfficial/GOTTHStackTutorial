package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	mux := http.NewServeMux()
	//TODO: USE GLOBS INSTEAD OF REGULAR PARSE FILES (parsefiles is stupidly expensive and is stupid i can do much better)
	//	statictemplates, err := template.ParseGlob("static/.html")
	//	slicedtemplates, err := template.ParseGlob("slices/.html")

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/index.html", "slices/genericfooter.html", "slices/genericheader.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	mux.Handle("/", fs)

	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		//yeah, parsing all these files is a mess and I can do better. But how?
		t, err := template.ParseFiles("static/index.html", "slices/genericfooter.html", "slices/genericheader.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/about.html", "slices/genericfooter.html", "slices/genericheader.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)

	})

	//this function is GENIUS!
	mux.HandleFunc("GET /showmewhygoth/{num}", func(w http.ResponseWriter, r *http.Request) {

		tutparam := r.PathValue("num")
		//fmt.Printf("ID IS : %s", tutparam)
		if tutparam == "" {
			http.Redirect(w, r, "/tutorials", http.StatusMovedPermanently)
		}
		if tutparam == "tailwindminified.js" {
			return
		}
		intparam, err := strconv.Atoi(tutparam)
		if err != nil {

			panic(err)
		}

		parsethisfile := "slices/whatsgotthdetailsslice.html"

		switch intparam {
		case 1:
			parsethisfile = "slices/whatsgotthdetailsslice.html"
		case 2:
			parsethisfile = "slices/whygotthdetailsslice.html"
		case 3:
			parsethisfile = "slices/whyopinionatedgotth.html"
		case 4:
			parsethisfile = "slices/whoisthisfor.html"
		case 5:
			parsethisfile = "slices/experiencewithaitools.html"
		case 6:
			parsethisfile = "slices/whynotdaisyui.html"
		case 10:
			parsethisfile = "slices/whynoframeworkslice.html"
		}
		t, err := template.ParseFiles(parsethisfile)
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, (t.Execute(w, nil)))
	})

	mux.HandleFunc("GET /whygotth", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/whygotth.html", "slices/genericfooter.html", "slices/genericheader.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	mux.HandleFunc("GET /tutorials", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/tutorials.html", "slices/genericfooter.html", "slices/genericheader.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	mux.HandleFunc("GET /0/{id}", func(w http.ResponseWriter, r *http.Request) {
		tutparam := r.PathValue("id")
		fmt.Printf("ID IS : %s", tutparam)
		if tutparam == "" {
			http.Redirect(w, r, "/tutorials", http.StatusMovedPermanently)
		}
		if tutparam == "tailwindminified.js" {
			return
		}
		intparam, err := strconv.Atoi(tutparam)
		if err != nil {

			fmt.Fprintf(w, "404 error.", http.StatusGone)
			return
		}
		parseThisFile := "static/tutorials.html"

		switch intparam {
		case 0:
			parseThisFile = "static/noviceintrotogo.html"
		case 1:
			parseThisFile = "static/noviceintrotohtmlcss.html"
		case 2:
			parseThisFile = "static/novicehtmx.html"
		case 3:
			parseThisFile = "static/novicenethttp.html"
		case 4:
			//parseThisFile = "static/novice"
			//TODO: ADD NOVICE TEMPLATING HTML FILE
		case 5:
			parseThisFile = "static/novicetailwind.html"
		}

		t, err := template.ParseFiles(parseThisFile, "slices/genericheader.html", "slices/genericfooter.html")
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		t.Execute(w, parseThisFile)

	})

	mux.HandleFunc("GET /tailwindminified.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		thisfile, err := os.ReadFile("/static/tailwindminified.js")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "", thisfile)

	})

	fmt.Printf("localhost:8080 + \n")
	http.ListenAndServe(":8080", mux)
}
