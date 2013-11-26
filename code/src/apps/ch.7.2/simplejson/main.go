// Shows how to use go-simplejson to decode json.
package main

import (
	"fmt"
	. "github.com/bitly/go-simplejson"
)

func main() {
	js, err := NewJson([]byte(`{
	    "test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	    }
	}`))
	if err != nil {
		panic( err )
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Printf("arr = %+v\n", arr)
	fmt.Printf("i = %+v\n", i)
	fmt.Printf("ms = %+v\n", ms)
}
