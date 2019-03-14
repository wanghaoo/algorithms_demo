package main

import (
  "fmt"
)

type Stack struct {
  First *StackNode
  Count int
  Size int
}

type StackNode struct {
  Value int
  Next *StackNode
}

func NewStack(size int) (*Stack) {
  var stack = new(Stack)
  stack.Size = size
  return stack
}

func (s *Stack) len()(int) {
  return s.Count
}

func (s *Stack)isFull()(bool) {
  return s.Size == s.Count
}

func (s *Stack) pull(value int) {
  if s.isFull() {
    return
  }
  var oldFirst = s.First
  var node = new(StackNode)
  node.Value = value
  node.Next = oldFirst
  s.Count += 1
  s.First = node
}

func (s *Stack) pop() (int) {
  var oldFirst = s.First
  var newFirst = s.First.Next
  s.First = newFirst
  s.Count -= 1
  return oldFirst.Value
}

func main() {
 var iArray = []int{1,2, 3,4,5,6,7}
 var stack = NewStack(5)
  for _, i := range iArray {
    stack.pull(i)
  }
  fmt.Println("stack count :", stack.Count)
  for stack.Count > 0{
    fmt.Println(stack.pop())
  }
}