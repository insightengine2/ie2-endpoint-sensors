package lib

import (
	"errors"
	"fmt"
	sensortypes "ie2-endpoint-sensors/types"
	"strings"
)

func HandleSensorData(ev sensortypes.SensorEvent) error {

	switch strings.ToLower(ev.SensorType) {
	case sensortypes.EAtmospheric:
		data, err := sensorEventToAtmos(ev)

		if err != nil {
			return err
		}

		ProcessAtmosphericData(*data)
	default:
		return fmt.Errorf("unidentified sensor: %s", ev.SensorType)
	}

	return nil
}

func sensorEventToAtmos(ev sensortypes.SensorEvent) (*sensortypes.SensorAtmos, error) {

	if _, exists := ev.SensorData["time"]; !exists {
		return nil, errors.New("key 'time' is missing from sensordata")
	}

	if _, exists := ev.SensorData["temp"]; !exists {
		return nil, errors.New("key 'temp' is missing from sensordata")
	}

	if _, exists := ev.SensorData["humidity"]; !exists {
		return nil, errors.New("key 'humidity' is missing from sensordata")
	}

	return &sensortypes.SensorAtmos{
		SensorType: ev.SensorType,
		SensorId:   ev.SensorId,
		SensorData: sensortypes.SDAtmos{
			Time:        ev.SensorData["time"],
			Temperature: ev.SensorData["temp"],
			Humidity:    ev.SensorData["humidity"],
		},
	}, nil
}
