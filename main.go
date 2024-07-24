// Module: github.com/Aamjit/GoChat-Go
package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/Aamjit/GoChat-Go/setupRouter"
)

func main() {
	fmt.Println("************************************************")
	fmt.Println("Go-Chat Chat App v0.01")
	fmt.Println("************************************************")
	setupRouter.SetupRouter()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("************************************************")
	fmt.Println("Using IP:", listener.Addr().(*net.TCPAddr).IP)
	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	fmt.Println("************************************************")

	panic(http.Serve(listener, nil))
}
