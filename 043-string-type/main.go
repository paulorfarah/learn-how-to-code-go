package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Println("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
