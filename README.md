# Go RabbitMQ Learning

This project demonstrates the basics of using RabbitMQ with the Go programming language. It includes essential operations such as connecting to RabbitMQ, publishing messages, and consuming messages from a queue. The application is containerized using Docker and Docker Compose.

## Description

The project consists of two main files:

1. **`main.go`**: The main application file responsible for setting up the connection to RabbitMQ, publishing messages, and consuming messages in a separate goroutine.
2. **`rabbitmq.go`**: Contains the implementation of the RabbitMQ service that manages connection, publishing, and consuming messages.

## Installation

Before running the application, ensure that Docker and Docker Compose are installed on your system. You can find installation instructions for Docker and Docker Compose on their [official websites](https://docs.docker.com/get-docker/) and [Docker Compose documentation](https://docs.docker.com/compose/install/).

### Cloning the Repository

```bash
https://github.com/rabboni171/rabbitmq-demo.git
docker-compose up --build

To stop and remove the containers, use:
docker-compose down



