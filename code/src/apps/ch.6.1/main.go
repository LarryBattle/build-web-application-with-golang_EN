// Example code for Chapter 6.1 from "Build Web Application with Golang"
// Purpose: Show how to read and write to a cookie on the server side
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)
const (
	PORT = "8080"
)
var t *template.Template = template.Must(template.ParseFiles("./index.gtpl"))

func updateCookie(c *http.Cookie) *http.Cookie {
	if c == nil {
		c = &http.Cookie{
			Name:    "visited",
			Path:    "/",
			Expires: time.Time{},
		}
	}
	i, _ := strconv.Atoi(c.Value)
	c.Value = fmt.Sprintf("%v", 1+i)
	return c
}
func index(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("visited")
	http.SetCookie(w, updateCookie(c))
	t.Execute(w, r.Cookies())
}
func showBlank(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Set the path for favicon to prevent duplicate cookie updates.
	http.HandleFunc("/favicon.ico", showBlank)
	http.HandleFunc("/", index)
	fmt.Println("Listing to localhost:", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	checkError(err)
}
