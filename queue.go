package main

import (
  "fmt"
  "github.com/pkg/errors"
)

type Queue struct {
  Head  int
  Tail  int
  Nodes []interface{}
}

func NewQueue(size int) (Queue) {
  queue := Queue{}
  queue.Head = 0
  queue.Tail = -1
  queue.Nodes = make([]interface{}, size)
  return queue
}

func (queue *Queue) IsEmpty() (bool) {
  return queue.Tail <= 0 || queue.Head == queue.Length()
}

func (queue *Queue) IsFull() (bool) {
  return queue.Tail == queue.Length() - 1
}

func (queue *Queue) Length() (int) {
  return len(queue.Nodes)
}

func (queue *Queue) Size() (int) {
  return queue.Tail - queue.Head + 1
}

func (queue *Queue) Push(value interface{}) (error) {
  if queue.IsFull() {
    return errors.New("queue is full")
  }
  queue.Tail += 1
  queue.Nodes[queue.Tail] = value
  return nil
}

func (queue *Queue) Pop() (interface{}, error) {
  if queue.IsEmpty() {
    return nil, errors.New("queue is empty")
  }
  value := queue.Nodes[queue.Head]
  queue.Nodes[queue.Head] = nil
  queue.Head += 1
  return value, nil
}

func main() {
  queue := NewQueue(10)
  v, err := queue.Pop()
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println("value: ", v)
  }
  for i := 1; i < 12; i ++ {
    err := queue.Push(i)
    if err != nil {
      fmt.Println(err.Error())
    }
  }
  fmt.Println("queue size: ", queue.Size())
  for i := 0; i <= 10; i ++ {
    v, err = queue.Pop()
    if err != nil {
      fmt.Println(err.Error())
    } else {
      fmt.Println("value: ", v)
    }
  }
  fmt.Println("queue size: ", queue.Size())
}
