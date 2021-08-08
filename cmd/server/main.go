package main

import (
	"github.com/Keda87/learn-go-rabbitmq/service/app"
	"github.com/Keda87/learn-go-rabbitmq/service/config"
)

func main() {
	conf := config.GetConfig()
	instance := app.New(&conf)

	instance.Start()
	defer instance.Stop()
}
