package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := "localhost:29092"
	topic := "orders"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer func(writer *kafka.Writer) {
		err := writer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)

	message := `{
  "order_uid": "a123bcf9d4e56test",
  "track_number": "WBILMNEWTRACK",
  "entry": "NEWL",
  "delivery": {
    "name": "Alice Johnson",
    "phone": "+1234567890",
    "zip": "100200",
    "city": "New City",
    "address": "123 Main Street",
    "region": "Central",
    "email": "alice@example.com"
  },
  "payment": {
    "transaction": "a123bcf9d4e56test",
    "request_id": "REQ123456",
    "currency": "EUR",
    "provider": "payfast",
    "amount": 2999,
    "payment_dt": 1685000000,
    "bank": "beta",
    "delivery_cost": 500,
    "goods_total": 2499,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1111111,
      "track_number": "WBILMNEWTRACK",
      "price": 1500,
      "rid": "item12345a",
      "name": "Lipstick",
      "sale": 10,
      "size": "M",
      "total_price": 1350,
      "nm_id": 1010101,
      "brand": "GlamBrand",
      "status": 201
    },
    {
      "chrt_id": 2222222,
      "track_number": "WBILMNEWTRACK",
      "price": 999,
      "rid": "item67890b",
      "name": "Mascara",
      "sale": 5,
      "size": "L",
      "total_price": 949,
      "nm_id": 2020202,
      "brand": "BeautyCo",
      "status": 202
    }
  ],
  "locale": "fr",
  "internal_signature": "sig123",
  "customer_id": "alice123",
  "delivery_service": "dhl",
  "shardkey": "7",
  "sm_id": 88,
  "date_created": "2023-05-25T12:30:00Z",
  "oof_shard": "2"
}
`

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("b563feb7b2b84b6test"),
			Value: []byte(message),
			Time:  time.Now(),
		},
	)

	if err != nil {
		log.Fatalf("failed to write message: %v", err)
	}

	log.Println("message sent to Kafka successfully")
}
