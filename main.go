// Module: github.com/Aamjit/GoChat-Go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Aamjit/GoChat-Go/setupRouter"
	"github.com/joho/godotenv"
)

func main() {

	// err = godotenv.Load(filepath.Join(pwd, "./.env"))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("************************************************")
	fmt.Println("Go-Chat Chat App v0.01")
	fmt.Println("************************************************")
	setupRouter.SetupRouter()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("************************************************")
	fmt.Println("Using IP:", listener.Addr().(*net.TCPAddr).IP)
	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	fmt.Println("************************************************")

	panic(http.Serve(listener, nil))
}
