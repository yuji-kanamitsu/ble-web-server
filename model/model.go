package model

type PostData struct {
	Key  string `json:"key"`
	Meta Meta
	Body []Body
}

type Meta struct {
	Area     int    `json:"area"`
	Type     int    `json:"type"`
	SensorID string `json:"sensor_id"`
	DataTime int    `json:"data_time"`
}

type Body struct {
	T   int     `json:"t"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
	Ble []Ble
}

type Ble struct {
	Addr string `json:"addr"`
	Rssi int    `json:"rssi"`
}

// SensorTable
type SensorTable struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	Key  string
	Meta Meta
	Body []Body
}
