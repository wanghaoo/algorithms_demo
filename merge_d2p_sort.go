package main

import (
  "fmt"
  "math"
)

func mergeD2pSort(array []int) {
  var len = len(array)
  var aux = make([]int, len, len)
  for i := 1; i < len; i *= 2 {
    for j := 0; j < len-i; j += i + i {
      var mid = i + j - 1
      var e = math.Min(float64(j+i+i-1), float64(len-1))
      merge2(array, aux, j, mid, int(e))
    }
  }
}

func merge2(array []int, aux []int, s, m, e int) {
  for k := s; k <= e; k++ {
    aux[k] = array[k]
  }
  i := s
  j := m + 1
  for k := s; k <= e; k++ {
    if i > m {
      array[k] = aux[j]
      j++
    } else if j > e {
      array[k] = aux[i]
      i ++
    } else if aux[j] < aux[i] {
      array[k] = aux[j]
      j ++
    } else {
      array[k] = aux[i]
      i ++
    }
  }
}

func main() {
  var array = []int{1, 54, 4, 67, 8, 4577, 34, 77, 35, 889, 567, 33, 5778, 799, 44, 56, 78, 443, 567, 4456}
  mergeD2pSort(array)
  fmt.Println(array)
}
