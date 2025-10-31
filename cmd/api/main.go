package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.HandleHealth)
	mux.HandleFunc("/suggest", server.HandleSuggest)

	mux.Handle("/", http.FileServer(http.Dir("web")))

	srv := &http.Server{
		Addr:    ":9696",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server running at: http://localhost:9696")
		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("[ERROR] error on init server: %v", err)
		}
	}()

	<-stop
	fmt.Printf("\n[INFO] closing server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("\n[INFO] Saving dataset updated...")
	if err := server.learner.Save(); err != nil {
		log.Println(err)
	} else {
		fmt.Println("[INFO] Dataset save on success!")
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[ERROR] could not shuting down server: %v", err)
	}

	fmt.Println(" Bye! Bye...")
}
