package main

import (
	"log"

	"github.com/mjthecoder65/image-processing-service/api"
	"github.com/mjthecoder65/image-processing-service/config"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	server, err := api.NewServer(config)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = server.Start()

	if err != nil {
		log.Fatal(err)
	}

}
