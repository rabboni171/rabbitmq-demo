package main

import (
	"fmt"

	"github.com/rabboni171/rabbitmq-demo/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ Learning")

	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hi there")
	if err != nil {
		return err
	}

	// Запускаем потребление сообщений в отдельной горутине
	go app.Rmq.Consume()

	select {}
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("error setting up our application")
		fmt.Println(err)
	}
}
