package main

import (
  "fmt"
  "strings"
)

func main() {
  var str = "jbpnbwwd"
  fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
  if len(s) <= 1 {
    return len(s)
  }
  var arr = []byte(s)
  var m = make(map[int][]byte)
  var iarr = make([]byte, 0)
  for i := 0; i < len(arr); i++ {
    for j := i ; j < len(arr); j++ {
      if !strings.Contains(string(iarr), string(arr[j])) {
        iarr = append(iarr, arr[j])
      } else {
        m[i] = iarr
        iarr = make([]byte, 0)
        break
      }
    }
  }
  m[-1] = iarr
  var max = 0
  for _, v := range m {
    if len(v) > max {
      max = len(v)
    }
  }
  return max
}