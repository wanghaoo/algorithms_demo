package main

import "fmt"

func main () {
  var arr = []int{4,3,2,6,8,1,11,56,3,45,7,2,5,7,9,3}
  for i, a := range arr {
    var min = a
    var innerIndex = i
    for n := i ; n < len(arr); n++ {
      if min > arr[n] {
        min = arr[n]
        innerIndex = n
      }
    }
    arr[i],arr[innerIndex] = arr[innerIndex],arr[i]
  }
  for _, a := range arr {
    fmt.Print(a, ",")
  }
}