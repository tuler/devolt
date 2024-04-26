package entity

type StationRepository interface {
	FindAllStations() ([]*Station, error)
}

type Station struct {
	ID        string                 `json:"_id"`
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	Params    map[string]interface{} `json:"params"`
}

func NewStation(id string, latitude float64, longitude float64, params map[string]interface{}) *Station {
	return &Station{
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
		Params:    params,
	}
}