package main

import "fmt"

func main() {
  var arr = []int{1, -2, 3, 5, -6, 10, -10, -8}
  var thisSum, maxSum int
  for _, a := range arr {
    thisSum += a
    if thisSum > maxSum {
      maxSum = thisSum
    } else if thisSum < 0 {
      thisSum = 0
    }
  }
  fmt.Println(maxSum)
}