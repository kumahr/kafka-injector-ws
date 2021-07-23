package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"os"
)

func postMessage(c *fiber.Ctx, producer *kafka.Producer) error {
	topic := c.Params("topic")
	key := c.Query("key")
	message := c.Body()
	delivery_chan := make(chan kafka.Event, 10000)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Key: []byte(key),
		Value: []byte(message)},
		delivery_chan,
	)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.SendStatus(201)
}

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/topics/:topic", func(c *fiber.Ctx) error {
		return postMessage(c, producer)
	})
	app.Listen(":8585")
}
