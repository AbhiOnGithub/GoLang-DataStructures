package main

// importing fmt package and container/heap
import (
	"container/heap"
	"fmt"
)

// integerHeap a type
type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// IntegerHeap method -swaps the element of index i to index j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i] //python style swapping in golang a,b = b,a
}

//IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapinterface interface{}) {
	*iheap = append(*iheap, heapinterface.(int))
}

//IntegerHeap method -pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

// main method
func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	heap.Push(intHeap, 5)
	heap.Push(intHeap, 4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("popping minimum item: %d \n", heap.Pop(intHeap))
	}
}
