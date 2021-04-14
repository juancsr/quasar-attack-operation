package models

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// NotherShip represents the position and message after decript
type MotherShip struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
