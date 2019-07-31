package main

import "fmt"

func main() {
  var array = []int{1, 54, 4, 67, 8, 4577, 34, 77, 35, 889, 567, 33, 5778, 799, 44, 56, 78, 443, 567, 4456}
  quickSort(array, 0, len(array) - 1)
  fmt.Println(array)
}

func quickSort(array []int, s, e int) {
  if e <= s {
    return
  }
  var j = partition(array, s, e)
  fmt.Println("j: ", j, "s", s, "e", e, "array: ", array)
  quickSort(array, s , j - 1)
  quickSort(array, j + 1, e)
}

func partition(array []int, s, e int) (int) {
  var cs1 = s
  var v = array[s]
  var cs2 = e + 1
  for {
    cs1 ++
    for ; array[cs1] < v; cs1 ++ {
      if cs1 == e {
        break
      }
    }
    cs2 --
    for ; array[cs2] > v; cs2 -- {
      if s == cs2 {
        break
      }
    }
    if cs1 >= cs2 {
      break
    }
    array[cs1], array[cs2] = array[cs2], array[cs1]
  }
  array[s], array[cs2] = array[cs2], array[s]
  return cs2
}