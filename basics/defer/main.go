package main

import "fmt"

func main() {
	defer fmt.Println("I was written first")

	fmt.Println("I was written second, but I will be executed first")

	defer fmt.Println("I was written Third, but I will be executed second")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("We have recovered")
		}
	}()
	panic("I am panicing")
}
