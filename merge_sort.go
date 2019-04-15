package main

import "fmt"

func merge(a []int, aux []int, s, m, e int) ([]int) {
  for k := s; k <= e; k++ {
    aux[k] = a[k]
  }
  
  // merge back to a[]
  i := s
  j := m + 1
  for k := s; k <= e; k++ {
    if i > m {
      j++
      a[k] = aux[j]
    } else if j > e {
      i++
      a[k] = aux[i]
    } else if aux[j] < aux[i] {
      j++
      a[k] = aux[j]
    } else {
      i++
      a[k] = aux[i]
    }
  }
  return a
}

func sort(array []int) {
  var aux = make([]int, len(array), len(array))
  sliptArraySort(array, aux, 0, len(array)-1)
}

func sliptArraySort(a []int, aux []int, s, e int) {
  if e <= s {
    return
  }
  mid := s + (e-s)/2
  sliptArraySort(a, aux, s, mid)
  sliptArraySort(a, aux, mid+1, e)
  merge(a, aux, s, mid, e)
}

func main() {
  var array = []int{1, 54, 4, 67, 8, 4577, 34, 77, 35, 889, 567, 33, 5778, 799, 44, 56, 78, 443, 567, 4456}
  sort(array)
  fmt.Println(array)
}
