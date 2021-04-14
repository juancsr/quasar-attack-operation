package main

import (
	"log"

	"github.com/juancsr/quasar-attack-operation/routers"
)

// @title Quasar Attack Operation API
// @version 1.0
// @description MELI Technical Challenge
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email juancsr@pm.me

// @license.name GNU GENERAL PUBLIC LICENSE V.30
// @license.url https://fsf.org

// @host quasar-attack-operation.herokuapp.com
// @BasePath /
func main() {
	log.Println("-- Launching app --")
	routers.Start()
}
