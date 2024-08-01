package main

import (
	"log"
	"net/http"
)

func style(w http.ResponseWriter, r *http.Request) {
	// Rendering master css from static folder
	http.ServeFile(w, r, "static/css/master.css")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func abtPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("", style)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", abtPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
