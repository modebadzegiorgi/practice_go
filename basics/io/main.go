package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func readWholefile(name string) (string, error) {
	// This calls os.Open and io.ReadAll() underneath
	f, err := os.ReadFile(name)

	if err != nil {
		return "", err
	}

	return string(f), nil
}

func writeFile(name string, content []byte) error {
	return os.WriteFile(name, content, 0644)
}

func readRemoteFile() {

	resp, err := http.Get("https://example.com")
	if err != nil {
		slog.Error("Could not read remote file", "err", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	writeFile("example.html", body)

}

func copyFile() {
	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC

	f, _ := os.OpenFile("example_copy.html", flags, 0644)

	resp, err := http.Get("https://example.com")
	if err != nil {
		slog.Error("Could not read remote file", "err", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	io.Copy(f, resp.Body)
}

func main() {

	logs, err := readWholefile("./logs.txt")

	if err != nil {
		slog.Error("Could not read file", "err", err)
		os.Exit(1)
	}

	fmt.Println(logs)

	err = writeFile("./hello.txt", []byte("hello world \n"))

	if err != nil {
		slog.Error("Could not write file", "err", err)
	}
	file, _ := os.OpenFile("hello.txt", os.O_RDONLY, 0644)
	// Stdin/ Stdout and stderr are files so instead of fmt.Println below also works

	io.Copy(os.Stdout, file)

	readRemoteFile()
	copyFile()

	// To read files as streams
	fs, _ := os.Open("./logs.txt")
	defer fs.Close()

	scanner := bufio.NewScanner(fs)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
