package main

import "fmt"

func main() {
	deq := NewDequeue()
	deq.PushFront(5)
	deq.PushBack(3)
	deq.PopBack()
	deq.Print()
}

type Node struct {
	Data int
	Next *Node
}

type Dequeue struct {
	begin *Node
	end   *Node
}

func NewDequeue() Dequeue {
	end := Node{
		Data: 0,
		Next: nil,
	}

	begin := Node{
		Data: 0,
		Next: &end,
	}
	return Dequeue{
		begin: &begin,
		end:   &end,
	}
}

func (deq *Dequeue) PushFront(data int) {
	node := &Node{
		Data: data,
		Next: deq.begin,
	}
	deq.begin = node
}

func (deq *Dequeue) PushBack(data int) {
	node := &Node{
		Data: data,
		Next: nil,
	}
	deq.end.Next = node
	deq.end = node
}

func (deq *Dequeue) PopFront() int {
	val := deq.begin.Data
	deq.begin = deq.begin.Next
	return val
}

func (deq *Dequeue) PopBack() int {
	val := deq.end.Data
	// finding prev node
	prevNode := findPrevNode(deq.begin)
	prevNode.Next = nil
	deq.end = prevNode
	return val
}

func findPrevNode(node *Node) *Node {
	result := Node{}
	for current := node; current != nil; current = current.Next {
		if current.Next.Next == nil {
			return current
		}
	}
	return &result
}

func (deq *Dequeue) Print() {
	for node := deq.begin; node != nil; node = node.Next {
		fmt.Println(node.Data)
	}
}
