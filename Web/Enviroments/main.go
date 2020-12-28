package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	Init()
}

//Init return name
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	name := os.Getenv("NAME")
	fmt.Println(name)
}
