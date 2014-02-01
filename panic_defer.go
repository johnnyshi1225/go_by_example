package main

import (
	. "fmt"
	"os"
)

func createFile(file string) *os.File {
	Println("Creating...")
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	Println("Writing...")
	Fprintln(f, "data")
}

func closeFile(f *os.File) {
	Println("Closing...")
	f.Close()
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
