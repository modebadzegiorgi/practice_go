package main

import "fmt"

const answer = 42

const (
	a = iota * 2
	b
	c
)

func main() {

	fmt.Println(answer + uint16(2))
	fmt.Println(answer + int32(2))

	fmt.Println(a, b, c)
}
