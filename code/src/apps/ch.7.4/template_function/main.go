package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	return strings.Replace(s, "@", " at ", -1)
}

func main() {
	t := template.New("example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(
		`Hi, I'm {{.UserName}}!
{{range .Emails}}
  Email: {{.|emailDeal}}
{{end}}
{{with .Friends}}
{{range .}}
  {{.Fname}} is my friend.
{{end}}
{{end}}
`)
	f1 := Friend{Fname: "Tom"}
	f2 := Friend{Fname: "Bob"}
	p := Person{
		UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
