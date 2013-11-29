package main

import (
    "os"
    "text/template"
)

func createAndExecuteTemplate(temp string){
	t := template.New("test")
	t = template.Must(t.Parse(temp))
    t.Execute(os.Stdout, nil)
}
func main() {
	// If the value of the pipeline is empty, no output is generated;
	// otherwise, the if block is executed.  The empty values are false, 0, any
	// nil pointer or interface value, and any array, slice, map, or
	// string of length zero. 
    createAndExecuteTemplate("Empty pipeline if: {{if ``}} NOT shown. {{end}}\n")
	createAndExecuteTemplate("Non-empty pipeline if demo: {{if `anything`}} IS shown. {{end}}\n")
	createAndExecuteTemplate("if-else demo: {{if `anything`}} if part {{else}} else part.{{end}}\n")
}