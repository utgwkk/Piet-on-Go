package main

func CalculateColor(r, g, b uint32) int {
  const (
    FF = 255
    CO = 192
    OO = 0
  )

  if r == FF {
    if g == FF {
      if b == FF { // white (#FFFFFF)
        return 18
      } else if b == CO { // light yellow (#FFFFC0)
        return 3
      } else { // yellow (#FFFF00)
        return 4
      }
    } else if g == CO {
      if b == FF { // light magenta (#FFC0FF)
        return 15
      }else if b == CO { // light red (#FFC0C0)
        return 0
      }
    } else {
      if b == FF { // magenta (#FF00FF)
        return 16
      } else if b == OO { // red (#FF0000)
        return 1
      }
    }
  } else if r == CO {
    if g == FF {
      if b == FF { // light cyan (#C0FFFF)
        return 9
      } else if b == CO { // light green (#C0FFC0)
        return 6
      }
    } else if g == CO {
      if b == FF { // light blue (#C0C0FF)
        return 12
      } else if b == OO { // dark yellow (#C0C000)
        return 5
      }
    } else {
      if b == CO { // dark magenta (#C000C0)
        return 17
      } else if b == OO { // dark red (#C00000)
        return 2
      }
    }
  } else {
    if g == FF {
      if b == FF { // cyan (#00FFFF)
        return 10
      } else if b == OO { // green (#00FF00)
        return 7
      }
    } else if g == CO {
      if b == CO { // dark cyan (#00C0C0)
        return 11
      } else if b == OO { // dark green (#00C000)
        return 8
      }
    } else {
      if b == FF { // blue (#0000FF)
        return 13
      } else if b == CO { // dark blue (#0000C0)
        return 14
      }
    }
  }

  // If the color is additional, it behaviors as black (#000000).
  return 19
}
