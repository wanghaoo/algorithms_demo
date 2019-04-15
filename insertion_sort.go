package main

import "fmt"

func insertion (array []int) (sortArray []int) {
  for i := 1; i < len(array); i ++ {
    for j := i; j > 0 && array[j] < array[j - 1]; j-- {
      array[j - 1], array[j] = array[j], array[j - 1]
    }
  }
  return array
}

func main() {
  var array = []int{1,4,56,7,3,7,34,7,44,2,865,544,7,55,3}
  sortArray := insertion(array)
  fmt.Println(sortArray)
}