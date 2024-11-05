package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "configs", "conf.yaml")

	data, _ := os.ReadFile(path)

	fmt.Println(string(data))

	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.Split(path))
	fmt.Println(os.TempDir())
	fmt.Println(filepath.Abs(path))
}
