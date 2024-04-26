package main

import (
	// "github.com/devolthq/devolt/internal/infra/ethereum"
	"github.com/devolthq/devolt/internal/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
)

func main() {
	msgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"session.timeout.ms": 6000,
		"group.id":           "devolt",
		"auto.offset.reset":  "latest",
	}

	// nodeUrl := os.Getenv("NODE_URL")
	// privateKey := os.Getenv("PRIVATE_KEY")
	// contractAddress := os.Getenv("CONTRACT_ADDRESS")
	// ethereumService, _ := ethereum.NewEthereumService(
	// 	nodeUrl,
	// 	privateKey,
	// 	contractAddress,
	// )

	kafkaRepository := kafka.NewKafkaConsumer(configMap, []string{os.Getenv("CONFLUENT_KAFKA_TOPIC_NAME")})
	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	for msg := range msgChan {
		// err := ethereumService.Insert(`{"station_id":"65fe7d0bfebaffb1eb337aaf","battery":354,"percent":35.38230884557721,"timestamp":"2024-03-24T08:39:17.13355113Z"}`)
		// if err != nil {
		// 	log.Printf("Error inserting data into Ethereum: %v", err)
		// 	continue
		// }
		// logs, _ := ethereumService.Get()
		// log.Printf("ETH Got parsed: %v", logs)

		log.Printf("Message on %s: %s 2", msg.TopicPartition, string(msg.Value))
	}
}
