package main

import "fmt"

func shell (array []int) ([]int) {
  length := len(array)
  h := 1
  for (h < length / 3) {
    h = h * 3 + 1
  }
  for h >= 1 {
    for i := h; i < length; i ++ {
      for j := i; j >= h && array[j] < array[j - h]; j -=h {
        array[j], array[j - h] = array[j - h], array[j]
      }
    }
    h /= 3
  }
  return array
}

func main() {
  var array = []int{54,7,78,3,347,92,1,78,3,6889,4,9,23,879,0,225}
  var sortArray = sort1(array)
  fmt.Println(sortArray)
}

func sort1(array []int) []int {
  l := len(array)
  h := 1
  for h < l / 5 {
    h = h * 5 + 1
  }
  for h >= 1 {
    for i := h; i < l; i ++ {
      for j := i; j >= h; j -= h {
        if array[j] > array[j - h] {
          array[j], array[j - h] = array[j - h], array[j]
        }
      }
    }
    h = h / 5
  }
  return array
}