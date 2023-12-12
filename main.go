package main

import (
	"fmt"

	"github.com/DeniesKresna/brigunaduty/config"
	"github.com/DeniesKresna/brigunaduty/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	cfg := config.CreateNewConfig()

	err = cfg.SetConfigApplication()
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	err = cfg.SetConfigDatabase()
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	err = service.Start(cfg)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
}
