package piet

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
)

func (p *Piet) None() { return }

func (p *Piet) Pop() { p.PopStack() }

func (p *Piet) Push() { p.PushStack(p.area) }

func (p *Piet) Add() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 { p.PushStack(b + a) }
}

func (p *Piet) Subtract() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 { p.PushStack(b - a) }
}

func (p *Piet) Multiply() {
  if len(p.stack) < 2 {
    return
  }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 {
    p.PushStack(b * a)
  }
}

func (p *Piet) Divide() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 {
    if a == 0 {
      return
    } else {
      p.PushStack(b / a)
    }
  }
}

func (p *Piet) Mod() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 {
    sgn := a < 0
    if a == 0 {
      p.PushStack(b)
      p.PushStack(a)
      return
    } else {
      v := int64(math.Abs(float64(b%a)))
      if sgn {
        p.PushStack(v * (-1))
      } else {
        p.PushStack(v)
      }
    }
  }
}

func (p *Piet) Not() {
  a, ok1 := p.PopStack()
  if ok1 { p.PushStack(BoolToInt64(a == 0)) }
}

func (p *Piet) Greater() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 { p.PushStack(BoolToInt64(a < b)) }
}

func (p *Piet) Pointer() {
  a, ok1 := p.PopStack()
  if ok1 {
    p.DP += int(a)
    p.DP %= 4
  }
}

func (p *Piet) Switch() {
  a, ok1 := p.PopStack()
  if ok1 {
    if a < 0 {
      a *= -1
      p.CC = true
    }
    for i := 0; i < int(a); i++ { p.CC = !p.CC }
  }
}

func (p *Piet) Duplicate() {
  a, ok1 := p.PopStack()
  if ok1 {
    p.PushStack(a)
    p.PushStack(a)
  }
}

func (p *Piet) Roll() {
  if len(p.stack) < 2 { return }
  a, ok1 := p.PopStack()
  b, ok2 := p.PopStack()
  if ok1 && ok2 {
    if b <= 0 {
      p.PushStack(b)
      p.PushStack(a)
      return
    }
    var reverse bool = false
    if a < 0 {
      a *= -1
      reverse = true
    }
    if int(b) > len(p.stack) {
      p.PushStack(b)
      p.PushStack(a)
      return
    }
    var x, xs []int64
    if int(b) == len(p.stack){
      x = []int64{}
      xs = []int64{}
    } else {
      x = []int64{p.stack[b]}
      xs = p.stack[b+1:]
    }
    for i := 0; i < int(a); i++ { p.stack = append(roll(p.stack[:b], reverse), append(x, xs...)...) }
  }
}

func (p *Piet) InNumber() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  a, err := strconv.Atoi(scanner.Text())
  if err != nil { return }
  p.PushStack(int64(a))
}

func (p *Piet) InChar() {
  var b []byte = make([]byte, 1)
  os.Stdin.Read(b)
  p.PushStack(int64(b[0]))
}

func (p *Piet) OutNumber() {
  a, ok1 := p.PopStack()
  if ok1 {
    fmt.Print(a)
    p.output += fmt.Sprintf("%d",a)
  }
}

func (p *Piet) OutChar() {
  a, ok1 := p.PopStack()
  if ok1 {
    fmt.Print(string(a))
    p.output += string(a)
  }
}