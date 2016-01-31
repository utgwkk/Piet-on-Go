package piet

import (
	"os"
  "strconv"
	"testing"
)

func TestAll(t *testing.T) {
	var testlist = [][3]string{
		{"./tests/alpha_filled.png", "abcdefghijklmnopqrstuvwxyz", "1"},
		{"./tests/Piet-1.gif", "Piet", "1"},
    {"./tests/gochiusa.png", "ああ〜心がぴょんぴょんするんじゃ〜", "10"}}

  os.Chdir(os.Getenv("GOPATH"))
	for _, tc := range testlist {
		var piet Piet
		reader, err := os.Open(tc[0])
    if err != nil {
      t.Errorf("%s", err)
      continue
    }
    //piet.EnableDebug()
    codel, _ := strconv.Atoi(tc[2])
		piet.SetCodelSize(codel)
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
