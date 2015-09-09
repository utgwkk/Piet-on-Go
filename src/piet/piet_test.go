package piet

import "testing"

func TestAll(t *testing.T) {
  var piet Piet
  var testlist = [][2]string{
    {"../../tests/alpha_filled.png", "abcdefghijklmnopqrstuvwxyz"},
    {"../../tests/Piet-1.gif", "Piet"}}

  for _, tc := range testlist {
    piet.New(tc[0])
    piet.debug = false
    piet.codelsize = 1
    piet.Run()
    actual := piet.output
    expected := tc[1]
    if actual != expected {
      t.Errorf("%v (actual) != %v (expected)", actual, expected)
    }
  }
}