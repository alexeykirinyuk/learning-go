package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
	"time"
)

var brockers = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}
var topicName = "order"

func newSyncProducer() (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brockers, cfg)

	return producer, err
}

func newAsyncProducer() (sarama.AsyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brockers, cfg)

	return producer, err
}

func prepareMessage(topic string, message []byte) *sarama.ProducerMessage {
	msg := sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(message),
		Partition: -1,
	}

	return &msg
}

type ProductInfo struct {
	SKU   int64
	Price float64
	Cnt   int64
}

type OrderInfo struct {
	UserId    int64
	CreatedAt time.Time
	Products  []ProductInfo
}

func generateOrder() OrderInfo {
	productCount := rand.Intn(20)

	products := make([]ProductInfo, 0, productCount)
	for i := 0; i < productCount; i++ {
		products = append(products, ProductInfo{
			SKU:   int64(rand.Intn(100)),
			Price: float64(rand.Intn(100)),
			Cnt:   int64(rand.Intn(10)),
		})
	}

	return OrderInfo{
		UserId:    int64(rand.Intn(100)),
		CreatedAt: time.Now().UTC(),
		Products:  products,
	}
}

func main() {
	syncProducer, err := newSyncProducer()
	if err != nil {
		log.Fatalf("newSyncProducer() returns err: %v\n", err)
	}

	asyncProducer, err := newAsyncProducer()
	if err != nil {
		log.Fatalf("newAsyncProducer() returns err: %v\n", err)
	}

	go func() {
		for err := range asyncProducer.Errors() {
			fmt.Printf("Msg async err: %v\n", err.Err)
		}
	}()

	go func() {
		for success := range asyncProducer.Successes() {
			log.Printf("Msg written in async. Partition: %d, Offset - %d\n", success.Partition, success.Offset)
		}
	}()

	for {
		order := generateOrder()

		oiJson, err := json.Marshal(order)
		if err != nil {
			log.Fatalf("json.Marshal(order) returns err: %v\n", err)
		}

		msg := prepareMessage(topicName, oiJson)

		if rand.Int()%2 == 0 {
			partition, offset, err := syncProducer.SendMessage(msg)
			if err != nil {
				fmt.Printf("Msg sync err: %v\n", err)
			} else {
				log.Printf("Msg written in sync. Partition: %d, Offset - %d\n", partition, offset)
			}
		} else {
			asyncProducer.Input() <- msg
		}

		time.Sleep(500 * time.Millisecond)
	}
}
