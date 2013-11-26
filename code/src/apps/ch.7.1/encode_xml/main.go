// Shows how to encode XML.
// Example code for Chapter 7.1 from "Build Web Application with Golang".
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}
func checkError(e error){
	if e != nil {
		panic( e )
	}
}
func main(){
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}