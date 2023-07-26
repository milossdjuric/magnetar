package main

import (
	"github.com/c12s/magnetar/configs"
	"github.com/c12s/magnetar/startup"
	"log"
)

func main() {
	config, err := configs.NewFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	err = startup.StartApp(config)
	if err != nil {
		log.Fatalln(err)
	}
}
