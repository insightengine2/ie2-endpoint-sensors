package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SensorEvent struct {
	// sample properties below
	SensorType string            `json:"sensortype"`
	SensorId   int16             `json:"sensorid"`
	SensorData map[string]string `json:"sensordata"`
}

func HandleRequest(context context.Context, ev events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "Success!",
	}

	data := SensorEvent{}
	err := json.Unmarshal([]byte(ev.Body), &data)

	if err != nil {

		res.StatusCode = 500
		res.Body = err.Error()

		return res, err
	}

	echo := fmt.Sprintf("Received data for SensorType %s and Id %d\r\n", data.SensorType, data.SensorId)

	for key, val := range data.SensorData {
		echo += fmt.Sprintf("Received Data %s with Value %s\r\n", key, val)
	}

	res.Body = echo

	// config.LoadDefaultConfig(context)

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
