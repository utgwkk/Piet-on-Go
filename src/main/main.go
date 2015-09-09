package main

import (
  "bufio"
  "fmt"
  "flag"
  "os"
  "os/exec"
  "strings"
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

  reader, err := os.Open(*fname)

  if err != nil {
    scanner := bufio.NewScanner(os.Stdin)
    var str string = ""
    for scanner.Scan() {
      str += scanner.Text()
    }
    altreader := strings.NewReader(str)
    pi.New(altreader)
  } else {
    pi.New(reader)
    defer reader.Close()
  }

  pi.Run()
  fmt.Println()
}
