package main

import (
  "piet"
  "bufio"
  "fmt"
  "flag"
  "os"
  "os/exec"
  "strings"
)

func main() {
  var pi piet.Piet
  var fname = flag.String("i", "filename", "source code file name")
  var debug = flag.Bool("debug", false, "if true, enables debug mode")
  var codelsize = flag.Int("codel", 1, "codel size of input")
  var limit = flag.Int("e", 100000, "the number of execution steps")
  flag.Parse()

  if os.DevNull == "/dev/null" { // UNIX
    // disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // restore the echoing state when exiting
    defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
  }

  if *debug { pi.EnableDebug() }
  pi.SetCodelSize(*codelsize)
  pi.SetExecLimit(*limit)

  reader, err := os.Open(*fname)

  if err != nil {
    scanner := bufio.NewScanner(os.Stdin)
    var str string = ""
    for scanner.Scan() { str += scanner.Text() }
    altreader := strings.NewReader(str)
    pi.New(altreader)
  } else {
    pi.New(reader)
    defer reader.Close()
  }

  pi.Run()
  fmt.Println()
}
