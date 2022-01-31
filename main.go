package main

import "fmt"

func main() {
	var f FileObject
	f, _ = OpenFile("sample.txt")
	defer f.Close()
	b, _ := ReadAll(f)
	fmt.Println(b)
}
