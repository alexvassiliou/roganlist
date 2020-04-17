package guest

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

const ratioOpeningText = "Average likes/dislikes ratio: "

// Guest represents a podcast guest
type Guest struct {
	Name  string
	Ratio string
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
			if a.Val == "guest-stats" {
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

func getRatio(z *html.Tokenizer) string {
	var result string
	tt := z.Next()
	if tt == html.StartTagToken {
		token := z.Token()
		for _, a := range token.Attr {
			if a.Val == "guest-stats-likes-ratio" {
			}
		}
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
