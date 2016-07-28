package main

import (
	"html/template"
	"net/http"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")
	t.Execute(w, r)
}

func main() {
	fs := http.FileServer(http.Dir("static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css", fs))
	http.HandleFunc("/", indexHandler)
	log.Println("Running server...")
	log.Println("Server listening at localhost:9000")
	http.ListenAndServe(":9000", nil)
}