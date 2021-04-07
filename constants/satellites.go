package constants

var Satellites map[string][2]float32

func init() {
	// Satelites
	_satellites := map[string][2]float32{
		"kenobi":    {-500.0, 200.0},
		"skywalker": {100.0, -100.0},
		"sato":      {500.0, 100.0},
	}
	Satellites = _satellites
}
