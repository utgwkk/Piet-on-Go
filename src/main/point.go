package main

import "fmt"

type Point struct {
  x int
  y int
}

func AddPoints(a, b Point) Point {
  return Point{a.x+b.x, a.y+b.y}
}

func (p Point) Equal(q Point) bool {
  return p.x == q.x && p.y == q.y
}

func (p Point) ToString() string {
  return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func Lesser(p, q Point, b bool) bool {
  if b { // compare x
    return p.x < q.x
  } else {
    return p.y < q.y
  }
}

func SortPointSlice(ps []Point, target, reverse bool) []Point {
  for i := 0; i < len(ps); i++ {
    for j := 1; j < len(ps)-i; j++ {
      if reverse {
        if Lesser(ps[j-1], ps[j], target) {
          ps[j], ps[j-1] = ps[j-1], ps[j]
        }
      } else {
        if Lesser(ps[j], ps[j-1], target) {
          ps[j], ps[j-1] = ps[j-1], ps[j]
        }
      }
    }
  }
  return ps
}