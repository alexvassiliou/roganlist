package guest

import "testing"
func TestTableParseRatio(t *testing.T) {

	var tests = []struct{
		input string
		expected float64
	} {
		{"Average likes/dislikes ratio: 8.2", 8.2},
		{"Average likes/dislikes ratio: 10.4", 10.4},
		{"Average likes/dislikes ratio: 1.25", 1.25},
		{"Average likes/dislikes ratio: 75.26", 75.26},
		{"Average likes/dislikes ratio: 10.00", 10.00},
	}

	for _, test := range tests {
		if output := parseRatio(test.input); output != test.expected {
			t.Errorf("expected was %v but actual shows %v", test.expected, output)
		}
	}
}