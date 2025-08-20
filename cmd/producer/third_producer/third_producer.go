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
  "order_uid": "b987xyz123456test",
  "track_number": "WBNEWTRACK987",
  "entry": "WEB",
  "delivery": {
    "name": "Bob Smith",
    "phone": "+9876543210",
    "zip": "300400",
    "city": "Oldtown",
    "address": "456 Elm Avenue",
    "region": "West",
    "email": "bob@example.com"
  },
  "payment": {
    "transaction": "b987xyz123456test",
    "request_id": "REQ987654",
    "currency": "USD",
    "provider": "stripe",
    "amount": 5499,
    "payment_dt": 1686000000,
    "bank": "gamma",
    "delivery_cost": 700,
    "goods_total": 4799,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 3333333,
      "track_number": "WBNEWTRACK987",
      "price": 1999,
      "rid": "item111aaa",
      "name": "Sneakers",
      "sale": 20,
      "size": "42",
      "total_price": 1599,
      "nm_id": 3030303,
      "brand": "SportyBrand",
      "status": 301
    },
    {
      "chrt_id": 4444444,
      "track_number": "WBNEWTRACK987",
      "price": 1200,
      "rid": "item222bbb",
      "name": "Backpack",
      "sale": 15,
      "size": "L",
      "total_price": 1020,
      "nm_id": 4040404,
      "brand": "UrbanPack",
      "status": 302
    },
    {
      "chrt_id": 5555555,
      "track_number": "WBNEWTRACK987",
      "price": 2600,
      "rid": "item333ccc",
      "name": "Headphones",
      "sale": 10,
      "size": "OneSize",
      "total_price": 2340,
      "nm_id": 5050505,
      "brand": "SoundMax",
      "status": 303
    }
  ],
  "locale": "en",
  "internal_signature": "sig987",
  "customer_id": "bob987",
  "delivery_service": "ups",
  "shardkey": "5",
  "sm_id": 99,
  "date_created": "2023-06-15T09:45:00Z",
  "oof_shard": "4"
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
