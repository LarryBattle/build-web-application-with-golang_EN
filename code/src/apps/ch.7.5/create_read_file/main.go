package main

import (
	"os"
)

var userFile string = "astaxie.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func createFile() {
	fout, err := os.Create(userFile)
	checkError(err)
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}
func readFile() {

	fl, err := os.Open(userFile)
	checkError(err)
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
func main() {
	createFile()
	readFile()
}
