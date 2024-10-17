package main

import (
	"encoding/json"
	"log"
)

type comment struct {
	Text string `form: "text" json: "text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comment", createComment)
	app.Listen(":3000")
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost: 8080"}
}

func createComment(c fiber.Ctx) error {
	cmt := new(comment)
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)
	c.JSON(&fiber.Map{
		"Success": true,
		"message": "comment pushed successfully!",
		"comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"Success": false,
			"Message": "comment not pushed",
		})
		return err
	}
	return err
}
