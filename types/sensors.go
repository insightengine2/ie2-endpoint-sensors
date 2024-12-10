package types

const (
	EAtmospheric string = "atmos"
)

// Receives sensor data from the body of the API request
type SensorEvent struct {
	// sample properties below
	SensorType string            `json:"sensortype"`
	SensorId   int16             `json:"sensorid"`
	SensorData map[string]string `json:"sensordata"`
}

// Humidity
type SDAtmos struct {
	Temperature string `json:"temperature"`
	Time        string `json:"time"`
	Humidity    string `json:"humidity"`
}

type SensorAtmos struct {
	SensorType string  `json:"sensortype"`
	SensorId   int16   `json:"sensorid"`
	SensorData SDAtmos `json:"sensordata"`
}
