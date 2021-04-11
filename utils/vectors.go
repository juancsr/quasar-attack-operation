package utils

import (
	"errors"
	"log"
	"math"

	"github.com/juancsr/quasar-attack-operation/constants"
	"github.com/juancsr/quasar-attack-operation/models"
)

// square of v number
func square(v float64) float64 {
	return v * v
}

// Normalice the vector with L2norm
func norm(p models.Circle) float64 {
	return math.Sqrt(square(p.X) + square(p.Y))
}

// Dot producto beetwen two vectors
func dot(p1, p2 models.Circle) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}

// Return a new Circle from the substract of p2 to p1
func subtract(p1, p2 models.Circle) models.Circle {
	return models.Circle{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
	}
}

// Return a new Circle from the sum of p2 to p1
func add(p1, p2 models.Circle) models.Circle {
	return models.Circle{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

// Return a new Circle from the division of p2 to p1
func divide(p models.Circle, v float64) models.Circle {
	return models.Circle{
		X: p.X / v,
		Y: p.Y / v,
	}
}

// Return a new Circle from the sum of p2 to p1
func multiply(p models.Circle, v float64) models.Circle {
	return models.Circle{
		X: p.X * v,
		Y: p.Y * v,
	}
}

// check if the Circle with coordinates x,y is inside of the circle p
func validateCircleInCircle(x, y float64, p models.Circle) bool {
	return float32(math.Pow(x-p.X, 2)+math.Pow(y-p.Y, 2)) == float32(math.Pow(p.R, 2))
}

// validate the solution checking all of the satellites agains the trilateration solution
func validateSolution(triPt models.Circle, circles ...models.Circle) (err error) {
	for _, c := range circles {
		if !validateCircleInCircle(triPt.X, triPt.Y, c) {
			err = errors.New("no solutions were found")
			return
		}
	}
	return
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	var (
		p1, p2, p3 = constants.Satellites["kenobi"], constants.Satellites["skywalker"], constants.Satellites["sato"]
	)

	p1.R = float64(distances[0])
	p2.R = float64(distances[1])
	p3.R = float64(distances[2])

	ex := divide(subtract(p2, p1), norm(subtract(p2, p1)))
	i := dot(ex, subtract(p3, p1))
	a := subtract(subtract(p3, p1), multiply(ex, i))
	ey := divide(a, norm(a))
	d := norm(subtract(p2, p1))
	j := dot(ey, subtract(p3, p1))

	_x := (square(p1.R) - square(p2.R) + square(d)) / (2 * d)
	_y := (square(p1.R)-square(p3.R)+square(i)+square(j))/(2*j) - (i/j)*_x

	trilaterationCircle := add(p1, add(multiply(ex, _x), multiply(ey, _y)))

	if err := validateSolution(trilaterationCircle, []models.Circle{p1, p2, p3}...); err != nil {
		log.Panic(err.Error())
	}

	x = float32(trilaterationCircle.X)
	y = float32(trilaterationCircle.Y)
	return
}
