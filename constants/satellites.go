package constants

import "github.com/juancsr/quasar-attack-operation/models"

var Satellites = map[string]models.Circle{
	"kenobi":    models.Circle{X: -500, Y: -200},
	"skywalker": models.Circle{X: 100, Y: -100},
	"sato":      models.Circle{X: 500, Y: 100},
}
