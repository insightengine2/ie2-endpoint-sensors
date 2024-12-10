package lib

import (
	"context"
	"encoding/json"
	sensortypes "ie2-endpoint-sensors/types"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

func ProcessAtmosphericData(data sensortypes.SensorAtmos) error {

	log.Print("Processing atmospheric data...")

	busname := os.Getenv("IE2_EVENTBUS_NAME")
	conf, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		log.Print(err)
		return err
	}

	if len(busname) <= 0 {
		log.Print("EventBus name is not defined. Falling back to default.")
		busname = "default"
	}

	log.Print("Marshaling data struct to json")
	event, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
		return err
	}

	client := eventbridge.NewFromConfig(conf)

	log.Printf("Pushing event to %s", busname)

	_, err = client.PutEvents(context.Background(), &eventbridge.PutEventsInput{
		Entries: []types.PutEventsRequestEntry{
			{
				EventBusName: aws.String("default"),
				Source:       aws.String("sensor.atmos"),
				DetailType:   aws.String("sensor.atmos.data"),
				Detail:       aws.String(string(event)),
			},
		},
	})

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
