package main

import (
	"log"
	"os"

	config "github.com/chutified/market-info/config"
	server "github.com/chutified/market-info/server"
)

func main() {
	log := log.New(os.Stdout, "[ERROR]", log.LstdFlags)

	// get config
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// server
	srv := server.New()
	err = srv.Set(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		log.Fatal(srv.Stop())
	}()
	// run
	log.Panic(srv.Start())
}
