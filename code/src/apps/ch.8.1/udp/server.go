// Shows how to communicate with a client using UDP connections.
// Example code for Chapter 8.1 from "Build Web Application with Golang".
package main

import (
    "fmt"
    "net"
	"strings"
    "os"
    "time"
)

func main() {
    service := ":1200"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)
    conn, err := net.ListenUDP("udp", udpAddr)
    checkError(err)
	fmt.Println("Server waiting for UDP connections to localhost" + service)
    for {
        handleClient(conn)
    }
}
func handleClient(conn *net.UDPConn) {
    var buf [512]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    if err != nil {
		fmt.Printf("Error: %v\n", err)
        return
    }
	fmt.Printf( "Received: %s\n", strings.Trim(string(buf[:]), "\x00" ))
    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime), addr)
	fmt.Printf("Sent: %s\n", daytime)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
        os.Exit(1)
    }
}
