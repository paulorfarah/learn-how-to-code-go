package main

import "fmt"

func main(){
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Printf("%v: %T\n", x, x)
	fmt.Println(y)
	fmt.Printf("%v: %T\n", y, y)
	fmt.Println(z)
	fmt.Printf("%v: %T\n", z, z)
}

