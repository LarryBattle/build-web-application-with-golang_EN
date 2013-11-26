// Shows how to decode XML.
// Example code for Chapter 7.1 from "Build Web Application with Golang".
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)
type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}
func checkError(e error){
	if e != nil {
		panic( e )
	}
}
func main(){
	file, err := os.Open("servers.xml")
	checkError(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	checkError(err)
	v := Servers{}
	checkError( xml.Unmarshal(data, &v) )
	fmt.Println(v.Description)
}