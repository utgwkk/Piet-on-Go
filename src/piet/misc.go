package piet

func all(l ...bool) bool {
  for _, b := range l {
    if !b { return false }
  }
  return true
}

func contains(l []Point, t Point) bool {
  for _, p := range l {
    if p.Equal(t) { return true }
  }
  return false
}

func BoolToInt64(b bool) int64 {
  if b {
    return 1
  } else {
    return 0
  }
}