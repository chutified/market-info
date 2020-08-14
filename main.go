package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	config "github.com/chutified/market-info/config"
	server "github.com/chutified/market-info/server"
	"github.com/gin-gonic/gin"
)

func main() {
	log := log.New(os.Stdout, "[STATUS] ", log.LstdFlags)

	// disable dubugging messages
	gin.SetMode(gin.ReleaseMode)

	// get config
	cfg, err := config.GetConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// server
	srv := server.New()
	err = srv.Set(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// gracefully shutting down
		err := srv.Stop()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Service stopped.")
	}()

	// control the server's termination
	errs := make(chan error)
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errs <- errors.New((<-sig).String())
	}()

	// run
	go func() {
		log.Printf("Server is running on the port %d.", cfg.APIPort)
		err := srv.Start()
		if err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	// catch the termination
	log.Printf("Service is being gracefully terminating (%v)\n", <-errs)
}
