package machinefi

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"github.com/machinefi/w3bstream/pkg/depends/protocol/eventpb"
	MQTT "github.com/machinefi/w3bstream/pkg/depends/conf/mqtt"
	"log"
)

type MachineFiMqttClient struct {
	Client *MQTT.Client
}

func NewMachineFiMqttClient(id string, broker *MQTT.Broker) (*MachineFiMqttClient, error) {
	client, err := broker.Client(id)
	if err != nil {
		return nil, err
	}
	return &MachineFiMqttClient{
		Client: client,
	}, nil
}

func (c *MachineFiMqttClient) Publish(topic string, payload []byte) error {
	event := &eventpb.Event{}
	err := proto.Unmarshal(payload, event)
	if err != nil {
			return err
	}
	jsonPayload, err := json.Marshal(event)
	if err != nil {
			return err
	}
	log.Printf("Publishing message %s to topic %s", string(jsonPayload), topic)
	return c.Client.WithTopic(topic).Publish(payload)
}