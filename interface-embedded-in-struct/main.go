package main

import "fmt"

type Fooer interface {
	Foo() string
}

type Container struct {
	Fooer
}

// sink takes a value implementing the Fooer interface.
func sink(f Fooer) {
	fmt.Println("sink:", f.Foo())
}

// TheRealFoo is a type that implements the Fooer interface.
type TheRealFoo struct {
}

func (trf TheRealFoo) Foo() string {
	return "TheRealFoo Foo"
}

func main() {
	co := Container{Fooer: TheRealFoo{}}

	sink(co)
}
