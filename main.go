package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tkak/echoapp/handlers"
	"github.com/tkak/echoapp/health"
)

const version = "1.0.0"

func main() {
	log.Println("Starting Echoapp...")

	httpAddr := os.Getenv("NOMAD_ADDR_http")
	if httpAddr == "" {
		log.Fatal("NOMAD_ADDR_http must be set and non-empty")
	}
	log.Printf("HTTP service listening on %s", httpAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.EchoHandler)
	mux.Handle("/version", handlers.VersionHandler(version))
	mux.HandleFunc("/healthz", health.HealthzHandler)
	mux.HandleFunc("/healthz/status", health.HealthzStatusHandler)

	httpServer := &http.Server{}
	httpServer.Addr = httpAddr
	httpServer.Handler = handlers.LoggingHandler(mux)

	log.Fatal(httpServer.ListenAndServe())
}
