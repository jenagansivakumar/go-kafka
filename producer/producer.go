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

func createComment(c fiber.Ctx) error {
	cmt := new(Comment)
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
		"message": "Comment pushed successfully!",
		"Comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"Success": false
			"Message": "Comment not pushed"
		})
	}
}
