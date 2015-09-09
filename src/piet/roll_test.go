package piet

import (
  "testing"
)

func TestRoll(t *testing.T) {
  var testcase = [][2][]int64{
    {{2,3,4,3,2,1}, {2,4,3,1}},
    {{-1,4,1,2,3,4}, {4,1,2,3}},
    {{4,6,1,2,3,4,5,6,7,8,9,10}, {5,6,1,2,3,4,7,8,9,10}}}
  for _, tc := range testcase {
    var a, b = tc[0][0], tc[0][1]
    var result = tc[0][2:]
    var x, xs []int64
    var reverse bool = false
    if a < 0 {
      a *= -1
      reverse = true
    }
    if int(b) == len(result){
      x = []int64{}
      xs = []int64{}
    } else {
      x = []int64{result[b]}
      xs = result[b+1:]
    }
    for i := 0; i < int(a); i++ { result = append(roll(result[:b], reverse), append(x, xs...)...) }
    if !equals(result, tc[1]) { t.Errorf("actual: %v, expected: %v", result, tc[1]) }
  }
}

func equals(p, q []int64) bool {
  if len(p) != len(q) { return false }
  for i := 0; i < len(p); i++ {
    if p[i] != q[i] { return false }
  }
  return true
}