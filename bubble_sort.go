package main

import "fmt"

func bubble_sort(array []int)([]int) {
  for i := 0; i < len(array); i++ {
  	for j := i + 1; j < len(array); j++ {
  		if array[i] < array[j] {
			array[j], array[i] = array[i], array[j]
		}
	}
  }
  return array
}

func main() {
	var array = []int{2,4,71,43,4,67,32,8812,34,445,445,772,4456,1,3,4,6}
	fmt.Println(bubble_sort(array))
}