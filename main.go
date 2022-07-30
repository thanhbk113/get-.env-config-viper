package main

import (
	"fmt"
	"log"
	"myapp/utils"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println("Value get from viper is:", config.DBDriver, config.DBSource, config.ServerAddress)
}
