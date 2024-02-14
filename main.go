package main

import (
	"log"

	"github.com/jrandyl/ranaria/server"
)

func main() {
	err := server.Start(":443")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
