package main

import "fmt"

func main() {
  var array = []int{1, 54, 4, 67, 8, 4577, 34, 77, 35, 889, 567, 33, 5778, 799, 44, 56, 78, 443, 567, 4456}
  sort(array)
  fmt.Println(array)
}

func sort(array []int) {
  var aux = make([]int, len(array), len(array))
  splitArraySort(array, aux, 0, len(array)-1)
}

func splitArraySort(a []int, aux []int, start, end int) {
  if end <= start {
    return
  }
  mid := start + (end - start) / 2
  splitArraySort(a, aux, start, mid)
  splitArraySort(a, aux, mid + 1, end)
  merge(a, aux, start, mid, end)
}

func merge(a []int, aux []int, s, m, e int) ([]int) {
  for k := s; k <= e; k++ {
    aux[k] = a[k]
  }
  i := s
  j := m + 1
  for k := s; k <= e; k++ {
    if i > m {
      a[k] = aux[j]
      j++
    } else if j > e {
      a[k] = aux[i]
      i++
    } else if aux[j] < aux[i] {
      a[k] = aux[j]
      j++
    } else {
      a[k] = aux[i]
      i++
    }
  }
  return a
}
