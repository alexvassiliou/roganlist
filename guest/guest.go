package guest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Guest represents a podcast guest
type Guest struct {
	Name  string  `json:"Name"`
	Ratio float64 `json:"Ratio"`
}

// Guests is a slice of guest
type Guests []Guest

func (g Guests) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: All guests endpoint")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(g)
}

// ParseHTML returns a slice of guests from the given html body
func ParseHTML(r io.Reader) []Guest {
	var result []Guest

	z := html.NewTokenizer(r)

	for {
		tokenType := z.Next()

		switch {
		case tokenType == html.ErrorToken:
			return result
		case tokenType == html.StartTagToken:
			token := z.Token()

			for _, a := range token.Attr {
				if strings.Contains(a.Val, "guests-item") {
					var g Guest
					g.getGuestAttributes(z)
					result = append(result, g)
				}
			}
		}
	}
	return result
}

func (g *Guest) getGuestAttributes(z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.StartTagToken {
		token := z.Token()
		for _, a := range token.Attr {
			if a.Val == "guest-name" {
				g.Name = getName(z)
				g.Ratio = getRatio(z)
			}
		}
	}
}

func getName(z *html.Tokenizer) string {
	var result string
	tt := z.Next()

	if tt == html.StartTagToken {
		result = extractText(z)
	}
	return result
}

func getRatio(z *html.Tokenizer) float64 {
	var result float64
	for {
		tokenType := z.Next()
		if tokenType == html.StartTagToken {
			for _, a := range z.Token().Attr {
				if a.Val == "guest-stats-likes-ratio" {
					result = parseRatio(extractText(z))
					return result
				}
			}
		}
	}
}

func extractText(z *html.Tokenizer) string {
	var result string
	tt := z.Next()
	if tt == html.TextToken {
		result = z.Token().Data
	}
	return result
}

func parseRatio(input string) float64 {
	re := regexp.MustCompile("\\D....")
	test := re.ReplaceAllString(input, "")
	i, err := strconv.ParseFloat(test, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// func (g Guests) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("templates/guests.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err2 := t.Execute(w, g)
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}
// }
