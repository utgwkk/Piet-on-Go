package piet

import (
  "fmt"
)

const (
  white = 18
  black = 19
)

func (p *Piet) Run() {
  var cont bool = true
  var dpstat = []string{"RIGHT", "DOWN", "LEFT", "UP"}
  var ccstat = []string{"LEFT", "RIGHT"}
  for i := 1; cont; i++ {
    if p.debug {
      fmt.Printf("step %04d : %s ; DP=%s, CC=%s\n", i, p.now.ToString(), dpstat[p.DP], ccstat[BoolToInt64(!p.CC)])
    }
    cont = p.Step()
    if p.debug {
      fmt.Println(p.stack)
    }
  }
  if p.debug {
    fmt.Println("Finished.")
  }
}

var dr = [...]Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, -1}}

var commandsstr = [][3]string{
  {"None", "Push", "Pop"},
  {"Add", "Subtract", "Multiply"},
  {"Divide", "Mod", "Not"},
  {"Greater", "Pointer", "Switch"},
  {"Duplicate", "Roll", "InNumber"},
  {"InChar", "OutNumber", "OutChar"}}

func (p *Piet) Step() bool {
  var commands = [][3]func(){
    {p.None, p.Push, p.Pop},
    {p.Add, p.Subtract, p.Multiply},
    {p.Divide, p.Mod, p.Not},
    {p.Greater, p.Pointer, p.Switch},
    {p.Duplicate, p.Roll, p.InNumber},
    {p.InChar, p.OutNumber, p.OutChar}}

  var edge, dest Point
  if p.GetCodel(p.now) == white {
    edge = p.now
    for a := AddPoints(edge, dr[p.DP]); p.IsMovableCodel(a) && p.GetCodel(a) == white; a = AddPoints(a, dr[p.DP]) {
      edge = a
    }
    dest = AddPoints(edge, dr[p.DP])
    if p.IsMovableCodel(dest) {
      p.from = edge
      p.now = dest
      return true
    }
  }
  var sc = []Point{}
  p.CalculateArea(&sc, p.GetCodel(p.now), p.now)
  p.area = int64(len(sc))

  // determine where to go
  for i := 0; i < 8; i++ {
    edge = p.GetEdge(sc)
    dest = AddPoints(edge, dr[p.DP])
    if p.IsMovableCodel(dest){
      break
    }
    if i % 2 == 0 {
      p.CC = !p.CC
    } else {
      p.DP = (p.DP + 1) % 4
    }
    if i == 7 {
      // end of the program
      return false
    }
  }

  if p.GetCodel(p.now) != white && p.GetCodel(dest) != white {
    dh := (p.GetCodel(dest)/3 - p.GetCodel(p.now)/3 + 30) % 6
    dl := (p.GetCodel(dest) - p.GetCodel(p.now) + 30) % 3
    if p.debug {
      fmt.Println(commandsstr[dh][dl])
    }
    commands[dh][dl]()
  }
  p.from = edge
  p.now = dest
  return true
}

func (p *Piet) GetEdge(sc []Point) Point {
  var rev bool = p.DP < 2
  var target bool = p.DP % 2 == 0
  sc = SortPointSlice(sc, target, rev)
  if (p.DP % 2 == 0) == p.CC {
    rev = !rev
  }

  var r int = 1

  for ; r < len(sc); r++ {
    if target {
      if sc[r].x != sc[0].x {
        break
      }
    } else {
      if sc[r].y != sc[0].y {
        break
      }
    }
  }
  target = !target

  xs := sc[r:]
  sc = append(SortPointSlice(sc[:r], target, rev), xs...)
  return sc[0]
}

func (p *Piet) CalculateArea(sc *[]Point, color int, offset Point) {
  *sc = append(*sc, offset)
  for _, d := range dr {
    target := AddPoints(offset, d)
    if p.IsMovableCodel(target) {
      if !contains(*sc, target) && color == p.GetCodel(target) {
        p.CalculateArea(sc, color, target)
      }
    }
  }
}

func (p *Piet) IsMovableCodel(point Point) bool {
  return all(
    point.x >= 0,
    point.x < p.width,
    point.y >= 0,
    point.y < p.height,
    p.GetCodel(point) != black)
}