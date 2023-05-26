package main

import (
	"context"
	"log"
	"ps-user/internal/adapter"
)

func main() {

	if err := adapter.Run(); err != nil {
		log.Fatal(context.Background(), "error running ps-user: %s", err.Error())
	}

}
