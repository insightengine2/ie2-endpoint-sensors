package main

import (
	"context"
	"encoding/json"
	"fmt"

	"ie2-endpoint-sensors/lib"
	sensortypes "ie2-endpoint-sensors/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(context context.Context, ev events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "Success!",
	}

	data := sensortypes.SensorEvent{}
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

	err = lib.HandleSensorData(data)

	if err != nil {

		res.StatusCode = 500
		res.Body = err.Error()

		return res, err
	}

	res.Body = echo

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
