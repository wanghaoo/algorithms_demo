package main

import "fmt"

type WeightedQuickUnion struct {
  Id    []int
  Size  []int
  count int
}

func (qf *WeightedQuickUnion) Instance(n int) {
  qf.Id = make([]int, n, n)
  qf.Size = make([]int, n, n)
  for i := 0; i < n; i ++ {
    qf.Id[i] = i
  }
  qf.count = n
}

func (qf *WeightedQuickUnion) Count() int {
  return qf.count
}

func (qf *WeightedQuickUnion) IsConnection(p int, q int) bool {
  return qf.Id[p] == qf.Id[q]
}

func (qf *WeightedQuickUnion) Find(p int) int {
  for p != qf.Id[p] {
    p = qf.Id[p]
  }
  return p
}

func (qf *WeightedQuickUnion) validNum(p int) bool {
  var n = len(qf.Id)
  if p < 0 || p > n {
    return false
  }
  return true
}

func (qf *WeightedQuickUnion) Union(p int, q int) {
  var qfp = qf.Find(p)
  var qfq = qf.Find(q)
  if qfq == qfp {
    return
  }
  if qf.Size[qfp] > qf.Size[qfp] {
    qf.Id[qfp] = qfq
    qf.Size[qfp] ++
  } else {
    qf.Size[qfq] ++
    qf.Id[qfq] = qfp
  }
  qf.count --
}

func main() {
  var qf = new(WeightedQuickUnion)
  qf.Instance(10)
  array := [][]int{{9, 0}, {3, 4}, {5, 8}, {7, 2}, {2, 1}, {5, 7}, {0, 3}, {4, 2}}
  for _, arr := range array {
    if !qf.validNum(arr[0]) || !qf.validNum(arr[1]) {
      continue
    }
    qf.Union(arr[0], arr[1])
  }
  for i, id := range qf.Id {
    fmt.Println(fmt.Sprintf("index :%d parant value is %d", i, id))
  }
  fmt.Println("components: ", qf.count)
  printTree(qf.Id, []int{})
  
}

func printTree(ids []int, pids []int) {
  circleIds := make([]int, 0)
  hasChild := false
  for _, pid := range pids {
    for _, cid := range ids {
      if pid == cid {
        hasChild = true
        break
      }
    }
  }
  if len(pids) <= 0 {
    hasChild = true
  }
  if !hasChild {
    return
  }
  if len(pids) <= 0 {
   for i, id := range ids {
     if i == id {
       fmt.Print("|", i, "|")
       circleIds = append(circleIds, i)
     } else {
       fmt.Print("| |")
     }
   }
  } else {
   for i , id := range ids {
     if i == id {
       fmt.Print("| |")
       continue
     }
     var index = -1
     for _, pid := range pids {
       if pid == id {
         index = i
         circleIds = append(circleIds, i)
       }
     }
     if index > -1 {
       fmt.Print("|", index, "|")
     } else {
       fmt.Print("| |")
     }
   }
  }
  fmt.Println("")
  if hasChild {
    printTree(ids, circleIds)
  }
}
