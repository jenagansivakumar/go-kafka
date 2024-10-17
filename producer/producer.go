package main

type comment struct {
	Text string `form: "text" json: "text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
}
