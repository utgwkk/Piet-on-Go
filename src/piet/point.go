package piet

import "fmt"

type Point struct {
	x int
	y int
}

func AddPoints(a, b Point) Point { return Point{a.x + b.x, a.y + b.y} }

func (p Point) Equal(q Point) bool { return p.x == q.x && p.y == q.y }

func (p Point) ToString() string { return fmt.Sprintf("(%d, %d)", p.x, p.y) }

func Lesser(p, q Point, b, reverse bool) bool {
	if b { // compare x
		if reverse {
			return p.x >= q.x
		} else {
			return p.x < q.x
		}
	} else {
		if reverse {
			return p.y >= q.y
		} else {
			return p.y < q.y
		}
	}
}

func SortPointSlice(ps []Point, target, reverse bool) []Point {
	if len(ps) <= 1 {
		return ps
	}
	var left = []Point{}
	var right = []Point{}
	x := ps[0]
	for _, d := range ps[1:] {
		if Lesser(d, x, target, reverse) {
			left = append(left, d)
		} else {
			right = append(right, d)
		}
	}
	return append(SortPointSlice(left, target, reverse), append([]Point{x}, SortPointSlice(right, target, reverse)...)...)
}
