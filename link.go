package main

import "fmt"

type Node struct {
  Value interface{}
  Next *Node
}

type Link struct {
  N int
  Item *Node
}

func (link *Link)Instance() {
  link.N = 0
  link.Item = &Node{}
}

func (link *Link)Push(value interface{}) {
  if link.N == 0 {
    link.Item = &Node{Value: value, Next: nil}
  } else {
    oldItem := link.Item
    item := &Node{Value: value, Next: oldItem}
    link.Item = item
  }
  link.N ++
}

func (link *Link) Pop() (interface{}) {
  if link.N == 0 {
    return nil
  }
  nextItem := link.Item.Next
  item := link.Item.Value
  link.Item = nextItem
  link.N --
  return item
}

func (link *Link) Size() (int) {
  return link.N
}

func main() {
  link := new(Link)
  link.Instance()
  link.Push("a")
  link.Push("b")
  link.Push("c")
  link.Push("d")
  link.Push("e")
  link.Push("f")
  link.Push("g")
  fmt.Println("-------------------")
  fmt.Println(link.Pop())
  fmt.Println(link.Pop())
  fmt.Println(link.Pop())
  fmt.Println("-------------------")
  fmt.Println(link.Size())
  for link.Size() > 0 {
    fmt.Println(link.Pop())
  }
}