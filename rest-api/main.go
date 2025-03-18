package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rest-api/handlers"
	"time"
)

func main() {
	err := startApp()
	if err != nil {
		panic(err)
	}
}

func startApp() error {
	api := http.Server{
		Addr:         ":8081",
		ReadTimeout:  500 * time.Second,
		WriteTimeout: 500 * time.Second,
		IdleTimeout:  500 * time.Second,
		Handler:      handlers.API(),
	}

	serverErr := make(chan error)
	go func() {
		serverErr <- api.ListenAndServe()
	}()
	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	// signal.Notify will notify the given channel when someone produces the given os signal
	signal.Notify(shutdown, os.Interrupt)

	select {
	// listening for errors that might happen during server startup, usually port is already being used
	case err := <-serverErr:
		return err
	case <-shutdown:
		fmt.Println("Gracefully shutting down server...")
		// creating a timer of 5sec for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners
		err := api.Shutdown(ctx)
		if err != nil {
			//close immediately closes all active net. Listeners and any connections in state
			// forceful close
			err := api.Close()
			if err != nil {
				return err
			}
		}

	}
	return nil
}
