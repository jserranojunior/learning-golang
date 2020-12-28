package main

import (
	routes "./router"
)

func main() {

	server := routes.Routes()

	server.Run(":8090")
}
