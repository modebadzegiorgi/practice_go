package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

const filterSize = 1000

var filter [filterSize]bool

func hash(value string) int {
	hash := md5.Sum([]byte(value))

	var val int

	for _, value := range hash {
		val += int(value)
	}

	return val % filterSize
}

func readText(path string) []string {
	file, err := os.Open(path)

	line := []string{}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	return line
}

func main() {

	items := readText("./items.txt")

	for _, item := range items {
		itemHash := hash(item)

		filter[itemHash] = true
	}

	// lets check if items is in filter

	itemToCheck := "Toothpaste"

	itemToCheckHash := hash(itemToCheck)

	fmt.Println(filter[itemToCheckHash])

}
