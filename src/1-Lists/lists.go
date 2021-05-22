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
	intList.PushBack(6)
	intList.PushBack(7)
	intList.PushBack(10)
	intList.PushFront(0)

	display(intList)

	println("Removed from Front :", intList.Front().Value.(int))

	intList.Remove(intList.Front())

	println("Removed from Back :", intList.Back().Value.(int))

	intList.Remove(intList.Front())

	display(intList)
}

func display(intList list.List) {
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(int), "\t")
	}
	println("\n Total Items ", intList.Len())
}
