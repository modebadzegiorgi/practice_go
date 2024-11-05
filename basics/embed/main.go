package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
)

//go:embed file.txt
var s string

//go:embed file.txt
//go:embed files/*
var content embed.FS

func main() {
	fmt.Println(s)
	fmt.Println(content.ReadFile("file.txt"))

	fs.WalkDir(content, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.IsDir())
		return nil
	})
}
