package main

import (
  "fmt"
  "flag"
)

func main() {
  var pi Piet
  var fname = flag.String("i", "filename", "source code file name")
  var debug = flag.Bool("debug", false, "if true, enables debug mode")
  flag.Parse()
  err := pi.New(*fname)
  if err != nil {
    fmt.Println(err)
    return
  }
  pi.debug = *debug
  pi.Run()
}
