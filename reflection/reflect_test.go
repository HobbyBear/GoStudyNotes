package reflection

import (
	"testing"
)

func TestPrintStrut(t *testing.T) {
	var p *Person
	RecursivePrintlnTagName(p)
}

func TestDisplay(t *testing.T) {
	var p *Person
	Display("xch", p)
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestDisplay2(t *testing.T) {
	stranagelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worring and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screeplay (Nomin)",
		},
		Sequel: nil,
	}
	Display("strangelove", stranagelove)
}
