package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Parthiba-Hazra/auth-go/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.CreateHandler(&handler.Config{
		E: router,
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize the server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v", srv.Addr)

	// Wait until any kill signal
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	// This context infrom the server to finish the request under 5 second
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down the server..")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
