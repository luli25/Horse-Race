package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"horse-race/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", api.RaceWebsocketHandler)

	// Servir frontend
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Ejecutar server en background
	go func() {
		log.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown con Ctrl+C
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt)

	<-stopSignal
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down:", err)
	}

	log.Println("Server stopped gracefully")
}
