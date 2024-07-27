package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Connect - establishes a connection to our RabbitMQ instance
// and declares the queue we are going to be using
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")

	var err error

	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ: %w", err)
	}

	fmt.Println("Successfully connected to RabbitMQ")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %w", err)
	}

	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue: %w", err)
	}

	return nil
}

// Publish - takes in a string message and publishes to a queue
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",          // exchange
		"TestQueue", // routing key (queue name)
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %w", err)
	}
	fmt.Println("Message published:", message)
	return nil
}

// Consume - consumes messages from our test queue
func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue", // queue
		"",          // consumer tag
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	fmt.Println("Waiting for messages...")
	for msg := range msgs {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
