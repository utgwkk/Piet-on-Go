package piet

import (
	"image"
	_ "image/gif"
	_ "image/png"
	"io"
)

type Piet struct {
	codel     [][]int
	codelsize int
	CC        bool
	DP        int
	width     int
	height    int
	area      int64
	stack     []int64
	from      Point
	now       Point
	debug     bool
	output    string
	execlimit int
}

func (p *Piet) SetCodelSize(size int) {
	if size <= 0 {
		p.codelsize = 1
	} else {
		p.codelsize = size
	}
}

func (p *Piet) SetExecLimit(limit int) {
	if limit < 0 {
		p.execlimit = 100000
	} else {
		p.execlimit = limit
	}
}

func (p *Piet) EnableDebug() { p.debug = true }

func (p *Piet) New(reader io.Reader) error {
	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bounds := m.Bounds()

	p.codel = make([][]int, bounds.Max.Y-bounds.Min.Y)
	for y := bounds.Min.Y; y < bounds.Max.Y; y += p.codelsize {
		for x := bounds.Min.X; x < bounds.Max.X; x += p.codelsize {
			r, g, b, _ := m.At(x, y).RGBA()
			p.codel[(y-bounds.Min.Y)/p.codelsize] = append(p.codel[(y-bounds.Min.Y)/p.codelsize], CalculateColor(r>>8, g>>8, b>>8))
		}
	}

	p.width = (bounds.Max.X - bounds.Min.X) / p.codelsize
	p.height = (bounds.Max.Y - bounds.Min.Y) / p.codelsize
	p.DP = 0
	p.CC = true
	p.stack = []int64{}
	p.from = Point{0, 0}
	p.now = Point{0, 0}
	p.output = ""
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
