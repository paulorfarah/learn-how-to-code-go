package main

import "fmt"

var y = 43
var z = "Shaken, not Stirred"
var a = `James said, "Shaken, not stirred"`

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
//	z :=  "Bond, James"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	foo()
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

func foo() {
	fmt.Println("z again: ", z)
}
