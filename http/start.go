package http

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Start is used to start an HTTP server.
func (s *server) Start() {
	// Initialize the server and allow it to gracefully shut down.
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"duration for which the server waits for existing connections to finish",
	)
	flag.Parse()

	s.addRoutes()

	srv := &http.Server{
		Addr:         "127.0.0.1:8083",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.r,
	}

	// Run the server in a goroutine so it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("failed to initialize server: %+v", err)
		}
	}()

	log.Printf("starting server on %s", srv.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
