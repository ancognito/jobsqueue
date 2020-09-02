package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/ancognito/jobsqueue/internal/api"
	"github.com/ancognito/jobsqueue/internal/jobs"
)

type ServerConfig struct {
	Port               int `envconfig:"PORT" default:"8080"`
	TerminationSeconds int `envconfig:"TERMINATION_SECONDS" default:"15"`
}

func main() {
	var cfg ServerConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("unable to process config: %v", err)
	}

	queue := jobs.New()

	v1 := api.NewV1Handler(queue)
	mux := http.NewServeMux()
	mux.Handle("/", v1)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	log.Printf("starting server on port %d\n", cfg.Port)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	log.Print("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TerminationSeconds)*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("error shutting down: %v", err)
	}
}
