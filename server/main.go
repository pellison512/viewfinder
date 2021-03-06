package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pellison512/viewfinder/server/data/v2"
	handlers "github.com/pellison512/viewfinder/server/handlers/v2"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	memDataSvc := data.NewMemStoreDataSvc()

	healthHanlder := handlers.NewHealthHandler()
	windowsHandler, err := handlers.NewWindowsHandler(memDataSvc)
	if err != nil {
		log.Fatalf("error creating windows handler with error '%s', shutting down", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/windows", windowsHandler.POSTWindowsHandler).Methods("POST")
	r.HandleFunc("/windows/{title}", windowsHandler.GETWindowsHandler).Methods("GET")
	r.HandleFunc("/windows/", windowsHandler.GETAllWindowsHandler).Methods("GET")
	r.HandleFunc("/health", healthHanlder.HealthGETHandler).Methods("GET")
	//	http.HandleFunc("/headers", handlers.HeadersHandler)
	//http.HandleFunc("/windows", handlers.WindowsHandler)
	//http.ListenAndServe(":8090", nil)

	srv := &http.Server{
		Addr: "0.0.0.0:8090",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      gorHandlers.LoggingHandler(os.Stdout, gorHandlers.CORS()(r)), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
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
