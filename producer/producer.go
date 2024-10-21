package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type comment struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen("0.0.0.0:3000")
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"kafka:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic (%s)/partition (%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {
	cmt := new(comment)
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request format",
		})
	}
	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error encoding comment",
		})
	}
	err = PushCommentToQueue("comments", cmtInBytes)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to push comment to queue",
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"message": "comment pushed successfully!",
		"comment": cmt,
	})
}
