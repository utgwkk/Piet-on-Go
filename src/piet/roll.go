package piet

func roll(xs []int64, reverse bool) []int64 {
  if reverse {
    x := xs[len(xs)-1]
    xs = xs[:len(xs)-1]
    return append([]int64{x}, xs...)
  } else {
    x := xs[0]
    xs = xs[1:]
    return append(xs, x)
  }
}