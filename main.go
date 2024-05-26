package main

import (
	"bobby/server"
	"bobby/util"
	"log"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = server.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
