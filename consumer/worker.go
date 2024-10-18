package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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

	fmt.Println("Consumer started")
	sigChain := make(chan os.Signal, 1)
	signal.Notify(sigChain, syscall.SIGINT, syscall.SIGTERM)
}
