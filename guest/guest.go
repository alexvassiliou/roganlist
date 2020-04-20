package guest

import (
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

const ratioOpeningText = "Average likes/dislikes ratio: "

// Guest represents a podcast guest
type Guest struct {
	Name  string
	Ratio string
}

// Guests is a slice of guest
type Guests []Guest

func (g Guests) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/guests.html")
	if err != nil {
		log.Fatal(err)
	}

	err2 := t.Execute(w, g)
	if err2 != nil {
		log.Fatal(err2)
	}
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
	switch {
	case tt == html.ErrorToken:
		return
	case tt == html.StartTagToken:
		token := z.Token()
		for _, a := range token.Attr {
			if a.Val == "guest-name" {
				g.Name = getName(z)
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

func extractText(z *html.Tokenizer) string {
	var result string
	tt := z.Next()
	if tt == html.TextToken {
		result = z.Token().Data
	}
	return result
}
