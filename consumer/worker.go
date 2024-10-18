package main

func main() {
	topic := "comments"

	worker, err := ConnectConsumer([]string{"localhost:29092"})

	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, OffsetOldest)
	if err != nil {
		panic(err)
	}
}
