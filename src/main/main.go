package main

import (
  "fmt"
  "flag"
)

func main() {
  var pi Piet
  var fname = flag.String("i", "filename", "source code file name")
  var debug = flag.Bool("debug", false, "if true, enables debug mode")
  var codelsize= flag.Int("codel", 1, "codel size of input")
  flag.Parse()
  pi.debug = *debug
  if *codelsize <= 0 {
    pi.codelsize = 1
  } else {
    pi.codelsize = *codelsize
  }
  err := pi.New(*fname)
  if err != nil {
    fmt.Println(err)
    return
  }
  pi.Run()
}
