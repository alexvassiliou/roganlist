package main

import (
	"flag"
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

	http.HandleFunc("/", redirect)

	http.Handle("/guests/", guestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/guests/", 301)
}
