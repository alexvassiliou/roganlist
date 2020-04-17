package main

import (
	"flag"
	"fmt"
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

	guests := guest.ParseHTML(resp.Body)

	for _, g := range guests {
		fmt.Println(g.Name)
		fmt.Println(g.Ratio)
	}

}
