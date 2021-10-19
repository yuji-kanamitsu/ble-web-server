package model

type SensingData struct {
	// ID        string  `json:"id"`
	Timestamp int     `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Sensing Table
type SensingTable struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	Timestamp int
	Latitude  float64
	Longitude float64
}
