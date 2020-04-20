package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/alexvassiliou/roganlist/guest"
)

func main() {
	url := flag.String("url", "https://www.jrepodcast.com/guests/", "url to parse")
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	var guestHandler guest.Guests
	guestHandler = guest.ParseHTML(resp.Body)

	http.HandleFunc("/", homeHandler)

	http.Handle("/guests/", guestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// get the names and likes ratio
// display it on a template
// order the page from most controversial to least
// add links to the videos

// html parsing, templates, go routines, http handlers
