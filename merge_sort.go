package main

import "fmt"

func main() {
  //var array = []int{1, 54, 4, 67, 8, 4577, 34, 77, 35, 889, 567, 33, 5778, 799, 44, 56, 78, 443, 567, 4456}
  var array = []int{1, 54, 4, 67}
  sort(array)
  //fmt.Println(array)
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
  fmt.Println("f mid: ", mid)
  splitArraySort(a, aux, mid + 1, end)
  fmt.Println("sec mid: ", mid + 1)
  //merge(a, aux, start, mid, end)
}

//var array = []int{1, 54, 4, 67}
func merge(a []int, aux []int, s, m, e int) ([]int) {
  //fmt.Println("s: ", s,"m: ",  m,"e: ", e)
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
      i++
      a[k] = aux[i]
    } else if aux[j] < aux[i] {
      a[k] = aux[j]
      j++
    } else {
      i++
      a[k] = aux[i]
    }
  }
  fmt.Println(a)
  return a
}
