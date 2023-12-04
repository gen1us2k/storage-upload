package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"storage/api"
	"storage/config"
	"time"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	s, err := api.New(c)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}