package main

import "fmt"

func main () {
  var arr = []int{4,3,2,6,8,1,11,56,3,45,7,2,5,7,9,3}
  for i, _ := range arr {
    var min = i
    for _, j := range arr {
      if min > j {
        min = j
      }
    }
    arr[i] = min
  }
  for _, a := range arr {
    fmt.Print(a)
  }
}