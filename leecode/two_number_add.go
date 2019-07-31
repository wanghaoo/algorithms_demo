package main

import (
  "fmt"
)

type ListNode struct {
  Val int
  Next *ListNode
 }

func main() {
  l1 := new(ListNode)
  l1.Val = 2
  l1.Next = &ListNode{Val:2, Next: &ListNode{Val:6}}
  l2 := new(ListNode)
  l2.Val = 5
  l2.Next = &ListNode{Val:8, Next: &ListNode{Val:3}}
  r := addTwoNumbers(l1, l2)
  fmt.Println(r.Val)
  for r.Next != nil {
    fmt.Println(r.Next.Val)
    r.Next = r.Next.Next
  }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var returnVal *ListNode
  var r = &ListNode{}
  var i = 0
  for l1 != nil || l2 != nil {
    var v1, v2 int
    if l1 != nil {
      v1 = l1.Val
    }
    if l2 != nil {
      v2 = l2.Val
    }
    val := v1 + v2
    decade := 0
    if val >= 10 {
      decade = 1
    }
    r.Next = &ListNode{Val:val % 10}
    if i == 0 {
     returnVal = r
     i ++
    }
    r = r.Next
    if l1 != nil {
      l1 = l1.Next
    }
    if l2 != nil {
      l2 = l2.Next
    }
    if decade > 0 {
      if l1 == nil {
        l1 = &ListNode{Val:decade}
      } else {
        l1.Val += decade
      }
    }
  }
  return returnVal.Next
}
