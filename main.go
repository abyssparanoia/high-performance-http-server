package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"golang.org/x/net/netutil"
)

func main() {

	var server *http.Server

	defer func() {
		err := recover()
		if err != nil {
			if e, ok := err.(error); ok {
				err = errors.Wrapf(e, "panic occurred outside of the request processing")
			}
			log.Fatalln(err)
		}
		gracefulShuttdown(server)
	}()

	addr := fmt.Sprintf(":%s", "8080")

	listner, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	routing(r)

	// server
	server = &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	quit := make(chan os.Signal, 1)
	log.Printf(fmt.Sprintf("[START] server. port: %s\n", addr))
	go func() {
		if err := server.Serve(netutil.LimitListener(listner, 1024)); err != http.ErrServerClosed {
			log.Fatalf("[CLOSED] server closed with error")
			quit <- syscall.SIGTERM
		}
	}()

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, os.Interrupt)
	log.Println(fmt.Sprintf("SIGNAL %d received, so server shutting down now...\n", <-quit))
}

func gracefulShuttdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if server != nil {
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalf("failed to gracefully shutdown")
		}
	}

	log.Fatalf("server shutdown completed")
}
