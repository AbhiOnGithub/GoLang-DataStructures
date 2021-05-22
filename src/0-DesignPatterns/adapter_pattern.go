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
	adapter.adaptee.convert()
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// main method
func main() {

	//this is how an interface is binding to class
	var processor IProcess = Adapter{}
	processor.process()
}
