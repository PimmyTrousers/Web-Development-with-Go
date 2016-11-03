package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var contactTemplate *template.Template
var faqTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactTemplate.Execute(w, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faqTemplate.Execute(w, nil)

}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "404 Bitches ")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layout/footer.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layout/footer.html")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)

}
