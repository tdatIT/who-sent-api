package main

import (
	"fmt"
	server "github.com/tdatIT/who-sent-api/internal"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Starting application...")
	defer log.Fatalf("[Info] Application has closed")

	serv, err := server.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		serv.Shutdown()
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serv.REST().Start(serv.Config().Server.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	wg.Wait()
}
