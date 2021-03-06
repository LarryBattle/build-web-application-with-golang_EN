package main
import (
    "net"
    "os"
    "fmt"
)
func main() {
    if len(os.Args) != 2 {
		filename := os.Args[0]
        fmt.Fprintf(os.Stderr, "Please pass an ip address.\n")
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", filename)
        os.Exit(1)
    }
    name := os.Args[1]
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        fmt.Println("The address is ", addr.String())
    }
    os.Exit(0)
}
