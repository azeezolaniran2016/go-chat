package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Info("Received request!")
		rw.Header().Add("type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg": "Welcome to go-chat app"}`))
		return
	}))

	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}

	go func() {
		log.Info("Starting server!")
		if err := srv.ListenAndServe(); err == nil {
			log.Fatalf("Error occurred while starting server. err: %+v", err)
			// return
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
