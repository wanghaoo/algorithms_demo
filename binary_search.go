package main

import "fmt"

func IndexOf (array []int, key int) (int) {
	var index = 0
	var len = len(array) - 1
	for index < len {
		var mid = index + (len - index) / 2
		if key < array[mid] {
			len = mid - 1
		} else if key > array[mid] {
			index = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main () {
	var array = []int{1,2,3,4,5,7,8,9,11,13}
	var index = IndexOf(array, 71)
	if index >= 0 {
		fmt.Println(fmt.Sprintf("the key at index %d of array", index))
	} else {
		fmt.Println("not found")
	}
}
