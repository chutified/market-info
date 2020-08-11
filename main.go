package main

import (
	"log"
	"os"

	config "github.com/chutified/market-info/config"
	server "github.com/chutified/market-info/server"
)

func main() {
	log := log.New(os.Stdout, "[STATUS] ", log.LstdFlags)

	// get config
	log.Printf("Getting config file...")
	cfg, err := config.GetConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// server
	log.Printf("Setting server...")
	srv := server.New()
	err = srv.Set(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		log.Printf("Stopping server...")
		log.Fatal(srv.Stop())
	}()
	// run
	log.Printf("Server is running on port %d.", cfg.APIPort)
	log.Panic(srv.Start())
}
