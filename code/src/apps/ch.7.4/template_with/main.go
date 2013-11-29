package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse(
		`hello {{.UserName}}!
{{range .Emails}}
  {{.}} is my email.
{{end}}
{{with .Friends}}
  {{range .}}
    {{.Fname}} is my friend.
  {{end}}
{{end}}
`)
	f1 := Friend{Fname: "Bob"}
	f2 := Friend{Fname: "Tom"}
	p := Person{
		UserName: "Astaxie",
		Emails:   []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
