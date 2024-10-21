
# go-kafka

**go-kafka** is a Go-based application that demonstrates the use of Apache Kafka for producing and consuming messages. It allows users to submit comments via an HTTP API, which are then published to a Kafka topic. The consumer listens for messages on the Kafka topic and processes them. The project is containerised using Docker and can be easily run locally using Docker Compose.


## Features

- **Go Fiber Framework**: Provides a lightweight HTTP server for handling API requests.
- **Kafka Producer & Consumer**: Implements both producer and consumer functionality using the Sarama library to communicate with Apache Kafka.
- **Docker & Docker Compose**: Ensures easy setup of the Kafka environment and Go application using Docker.
- **Graceful Shutdown**: The consumer listens for system signals to handle graceful shutdowns.
- **Scalable Design**: This project can be extended to scale with additional consumers, topics, and more complex processing.

## Prerequisites

- Docker & Docker Compose installed
- Go 1.20+ installed (if running without Docker)
- Apache Kafka (setup is included in Docker Compose)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/jenagansivakumar/go-kafka.git
   cd go-kafka 
2. Build and run the containers:

```bash
docker compose up -d
```
3. The API will be running at http://localhost:3000/api/v1/comments, and Kafka 
will be running on localhost:29092.
## API Reference

### Post message

```http
POST /api/v1/comments
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `text` | `string` |**Required**. The comment content|

**Example request**
```
curl -X POST http://localhost:3000/api/v1/comments \
     -H "Content-Type: application/json" \
     -d '{"text":"This is a sample comment"}'

```

## Project Structure

- **producer/producer.go**: Implements the HTTP API and Kafka producer logic.
- **worker/worker.go**: Implements the Kafka consumer that reads messages from the "comments" topic and processes them.
- **Dockerfile**: Defines the Docker image for building and running the consumer.
- **docker-compose.yml**: Defines the services for running Kafka, Zookeeper, and the Go application.
- **go.mod & go.sum**: Manage the project dependencies.
