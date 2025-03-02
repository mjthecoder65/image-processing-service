package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mjthecoder65/image-processing-service/api"
	"github.com/mjthecoder65/image-processing-service/config"
)

func main() {
	ctx := context.Background()

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := pgx.Connect(context.Background(), config.DatabaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(ctx)

	server, err := api.NewServer(config, conn)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
