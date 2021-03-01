package main

import (
	"log"
	"nacos-web/config"
	"nacos-web/nacos"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	nacos.Start()
}
