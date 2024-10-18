package main

func main() {
	topic := "comments"

	worker, err := ConnectConsumer([]string{"localhost:29092"})

	if err != nil {
		panic(err)
	}

}
