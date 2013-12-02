package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	path1 := "bareFolder"
	path2 := "folder1/folder2/folder3"

	fmt.Printf("Removing then creating folder `%s`.\n", path1)
	os.Remove(path1)
	checkError(os.Mkdir(path1, 0777))

	fmt.Printf("Removing then creating path `%s`.\n", path2)
	os.RemoveAll(path2)
	checkError(os.MkdirAll(path2, 0777))
}
