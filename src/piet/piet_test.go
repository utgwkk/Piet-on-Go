package piet

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	var testlist = [][2]string{
		{"../../tests/alpha_filled.png", "abcdefghijklmnopqrstuvwxyz"},
		{"../../tests/Piet-1.gif", "Piet"}}

	for _, tc := range testlist {
		var piet Piet
		reader, _ := os.Open(tc[0])
		piet.SetCodelSize(1)
		piet.New(reader)
		piet.Run()
		actual := piet.output
		expected := tc[1]
		if actual != expected {
			t.Errorf("%v (actual) != %v (expected)", actual, expected)
		}
	}
}
