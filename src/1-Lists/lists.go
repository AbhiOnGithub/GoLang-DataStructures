package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var intList list.List

	intList.PushBack(5)
	intList.PushBack(10)
	intList.PushFront(15)

	display(intList)

	intList.Remove(intList.Front())

	display(intList)
}

func display(intList list.List) {
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(int), "\t")
	}
	println("\n Total Items ", intList.Len())
}
