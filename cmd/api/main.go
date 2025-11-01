package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

import _ "net/http/pprof"

func main() {
	server := NewServer()
	printMemUsage()

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
		log.Println("pprof active on :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

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

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("=== GO Memory ===\n")
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
	fmt.Println("=========================")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
