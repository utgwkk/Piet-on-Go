package piet

func (p *Piet) PushStack(i int64) {
	p.stack = append([]int64{i}, p.stack...)
}

func (p *Piet) PopStack() (int64, bool) {
	if len(p.stack) == 0 {
		return 0, false
	} else {
		value := p.stack[0]
		p.stack = p.stack[1:]
		return value, true
	}
}
