package main

import (
	"context"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devolthq/devolt/internal/infra/kafka"
	"github.com/devolthq/devolt/pkg/cartesi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	msgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"session.timeout.ms": 6000,
		"group.id":           "devolt",
		"auto.offset.reset":  "latest",
	}

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to blockchain: %v", err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	instance, err := cartesi.NewInputBox(common.HexToAddress(os.Getenv("INPUT_BOX_CONTRACT_ADDRESS")), client)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	kafkaRepository := kafka.NewKafkaConsumer(configMap, []string{os.Getenv("CONFLUENT_KAFKA_TOPIC_NAME")})

	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()

	for msg := range msgChan {
		wg.Add(1)
		go func(msg *ckafka.Message) {
			defer wg.Done()
			if transaction, err := instance.AddInput(opts, common.HexToAddress(os.Getenv("APPLICATION_CONTRACT_ADDRESS")), msg.Value); err != nil {
				log.Printf("Failed to add input: %v", err)
			} else {
				log.Printf("Transaction sent with hash: %v and payload: %v", transaction.Hash().Hex(), msg.Value)
			}
		}(msg)
	}
	wg.Wait()
}
