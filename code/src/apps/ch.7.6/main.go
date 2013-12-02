package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func join() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	//Output:foo, bar, baz
}
func indexOf() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	//Output:4
	//-1
}
func repeat() {
	fmt.Println("ba" + strings.Repeat("na", 2))
	//Output:banana
}
func replace() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//Output:oinky oinky oink
	//moo moo moo
}
func split() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//Output:["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]
}
func trim() {
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	// Output:["Achtung"]
}
func fields() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//Output:Fields are: ["foo" "bar" "baz"]
}
func contains() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}
func append() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}
func format() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}

func parseStringToInt() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	fmt.Println(a, b, c, d)
}
func main() {
	parseStringToInt()
	contains()
	join()
	indexOf()
	repeat()
	replace()
	append()
	contains()
	fields()
	split()
	trim()
}
