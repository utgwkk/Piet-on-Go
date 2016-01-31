package piet

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	var testlist = [][2]string{
		{"./tests/alpha_filled.png", "abcdefghijklmnopqrstuvwxyz"},
		{"./tests/Piet-1.gif", "Piet"}}

  os.Chdir(os.Getenv("GOPATH"))
	for _, tc := range testlist {
		var piet Piet
		reader, err := os.Open(tc[0])
    if err != nil {
      t.Errorf("%s", err)
      continue
    }
    piet.EnableDebug()
		piet.SetCodelSize(1)
    piet.SetExecLimit(100000)
		piet.New(reader)
		piet.Run()
		actual := piet.GetOutput()
		expected := tc[1]
		if actual != expected {
			t.Errorf("%v (actual) != %v (expected)", actual, expected)
		}
	}
}
