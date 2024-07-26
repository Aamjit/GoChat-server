// Module: github.com/Aamjit/GoChat-Go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Aamjit/GoChat-Go/setupRouter"
	"github.com/joho/godotenv"
)

func main() {

	// err = godotenv.Load(filepath.Join(pwd, "./.env"))
	app_env := os.Getenv("APP_ENV")
	fmt.Println(app_env)

	if app_env == "DEV" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else if app_env == "PRD" {
		godotenv.Load()
	}

	envs := os.Environ()
	for _, e := range envs {
		fmt.Println(e)
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
