package main

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

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

type Consumer struct {
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (c *Consumer) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var order OrderInfo
		err := json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Printf("error when json unmasrshall: %v\n", err)
		}

		log.Printf("msg: %v\n", order)

		s.MarkMessage(msg, "")
	}

	return nil
}

func subscribe(ctx context.Context, topic string, cg sarama.ConsumerGroup) error {
	consumer := Consumer{}

	go func() {
		for {
			if err := cg.Consume(ctx, []string{topic}, &consumer); err != nil {
				log.Printf("err when cg.Consume(): %v\n", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}

var brockers = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

const topicName = "order"

func StartConsuming(ctx context.Context) error {
	cfg := sarama.NewConfig()
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.BalanceStrategyRoundRobin,
	}
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerGroup, err := sarama.NewConsumerGroup(brockers, "analytic_group", cfg)
	if err != nil {
		return err
	}

	if err = subscribe(ctx, topicName, consumerGroup); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := StartConsuming(ctx); err != nil {
		log.Fatalf("StartConsuming(ctx): %v", err)
	}

	for {
	}
}
