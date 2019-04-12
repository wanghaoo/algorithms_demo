package main

import "fmt"

type QuickFind struct {
  Id    []int
  count int
}

func (qf *QuickFind) Instance(n int) {
  qf.Id = make([]int, n, n)
  for i := 0; i < n; i ++ {
    qf.Id[i] = i
  }
  qf.count = n
}

func (qf *QuickFind) Count() int {
  return qf.count
}

func (qf *QuickFind) IsConnection(p int, q int) bool {
  return qf.Id[p] == qf.Id[q]
}

func (qf *QuickFind) Find(p int) int {
  return qf.Id[p]
}

func (qf *QuickFind) validNum(p int) bool {
  var n = len(qf.Id)
  if p < 0 || p > n {
    return false
  }
  return true
}

func (qf *QuickFind) Union(p int, q int) {
  var qfp = qf.Find(p)
  var qfq = qf.Find(q)
  if qfq == qfp {
    return
  }
  for i, id := range qf.Id {
    if id == qfp {
      qf.Id[i] = qfq
    }
  }
  qf.count --
}

func main() {
  var qf = new(QuickFind)
  qf.Instance(10)
  array := [][]int {{9,0},{3,4},{5,8},{7,2},{2,1},{5,7},{0,3},{4,2}}
  for _, arr := range array {
   if !qf.validNum(arr[0]) || !qf.validNum(arr[1]) {
     continue
   }
   qf.Union(arr[0], arr[1])
  }
  for i, id := range qf.Id {
   fmt.Println(fmt.Sprintf("index :%d parant value is %d", i , id))
  }
  fmt.Println(qf.count)
}