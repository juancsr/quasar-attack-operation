package main

import (
	"log"

	"github.com/juancsr/quasar-attack-operation/routers"
)

func main() {
	log.Println("-- Launching app --")
	routers.Start()
}
