// Example code for Chapter 6.3 from "Build Web Application with Golang"
// Purpose: Show how to create and handle sessions
// Goal: Create a simple login page that shows the amount of times visited.
// This method stores the session data in memory
package main

import (
	"fmt"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"net/http"
)

var (
	globalSessions *session.Manager
	t              *template.Template = template.Must(template.ParseFiles("./index.gtpl"))
)

type userInfo struct {
	AmountVisited int
	Username      string
}

const (
	PORT = "8080"
)

func (u *userInfo) increaseVisit() {
	if u != nil {
		u.AmountVisited += 1
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	var data userInfo
	if x := sess.Get("userInfo"); x != nil {
		data = x.(userInfo)
	}
	data.increaseVisit()
	sess.Set("userInfo", data)
	t.Execute(w, data)
}
func showBlank(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sess := globalSessions.SessionStart(w, r)
		r.ParseForm()
		sess.Set("userInfo", userInfo{Username: r.FormValue("username")})
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}
func main() {
	// Set the path for favicon to prevent duplicate cookie updates.
	http.HandleFunc("/favicon.ico", showBlank)
	http.HandleFunc("/", index)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	fmt.Println("Listening to localhost:", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	checkError(err)
}
