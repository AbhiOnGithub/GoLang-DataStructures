package main

// importing fmt package
import (
	"fmt"
)

//IProcess interface
type IProcess interface {
	process()
}

//Adaptee Struct
type Adaptee struct {
	adapterType int
}

//Adapter struct
type Adapter struct {
	adaptee Adaptee
}

//Adapter class method process
func (adapter Adapter) process() {
	fmt.Println("Adapter process")

	var val = adapter.adaptee.convert()

	fmt.Println(val)
}

// Adaptee class method convert
func (adaptee Adaptee) convert() (adapterType int) {
	fmt.Println("Adaptee convert method")
	adaptee.adapterType = 10
	return adaptee.adapterType
}

// main method
func main() {

	//this is how an interface is binding to class
	var processor IProcess = Adapter{}
	processor.process()
}
