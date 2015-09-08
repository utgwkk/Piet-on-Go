package main

import (
  "os"
  "image"
  _ "image/gif"
  _ "image/png"
)

type Piet struct {
  codel [][]int
  CC bool
  DP int
  width int
  height int
  area int64
  stack []int64
  from Point
  now Point
  debug bool
}

func (p *Piet) New(filename string) error {
  reader, err := os.Open(filename)

  if err != nil {
    return err
  }
  defer reader.Close()

  return p.NewFromFile(reader)
}

func (p *Piet) NewFromFile(reader *os.File) error {
  m, _, err := image.Decode(reader)
  if err != nil {
    return err
  }
  bounds := m.Bounds()

  p.codel = make([][]int, bounds.Max.Y - bounds.Min.Y)
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, _ := m.At(x, y).RGBA()
      p.codel[y-bounds.Min.Y] = append(p.codel[y-bounds.Min.Y], CalculateColor(r>>8, g>>8, b>>8))
    }
  }

  p.width = bounds.Max.X - bounds.Min.X
  p.height = bounds.Max.Y - bounds.Min.Y
  p.DP = 0
  p.CC = true
  p.stack = []int64{}
  p.from = Point{0, 0}
  p.now = Point{0, 0}
  return nil
}

func (p *Piet) GetCodel(point Point) int {
  if all(
    point.x >= 0,
    point.x < p.width,
    point.y >= 0,
    point.y < p.height) {
    return p.codel[point.y][point.x]
  } else {
    return black
  }
}