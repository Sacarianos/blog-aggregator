package main

import (
	"fmt"
	"log"

	"github.com/Sacarianos/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("Jaime")
	if err != nil {
		log.Fatalf("Failed to set user: %v", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read updated config: %v", err)
	}

	fmt.Println("Updated Config:", updatedCfg)

}
