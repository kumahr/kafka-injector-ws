package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/Shopify/sarama.v1"
)

func postMessage(c *fiber.Ctx, writer sarama.SyncProducer) error {
	topic := c.Params("topic")
	key := c.Query("key")
	message := c.Body()
	event := &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(message)}
	_, _, err := writer.SendMessage(event)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.SendStatus(201)
}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/topics/:topic", func(c *fiber.Ctx) error {
		return postMessage(c, producer)
	})
	app.Listen(":8585")
}
