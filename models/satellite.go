package models

type Satellite struct {
	Name     string   `json:"name" binding:"required"`
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}
