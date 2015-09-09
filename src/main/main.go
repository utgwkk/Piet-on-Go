package main

import (
  "fmt"
  "flag"
  "os"
  "os/exec"
)

func main() {
  var pi Piet
  var fname = flag.String("i", "filename", "source code file name")
  var debug = flag.Bool("debug", false, "if true, enables debug mode")
  var codelsize= flag.Int("codel", 1, "codel size of input")
  flag.Parse()

  if os.DevNull == "/dev/null" { // UNIX
    // disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // restore the echoing state when exiting
    defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
  }

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
  fmt.Println()
}
