package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func main() {
	flag.Parse()

	r := routes()

	s := &fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		log.Printf("starting server at %q", *addr)
		if err := s.ListenAndServe(*addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	log.Println("shutting down server")
	if err := s.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("shutting down server: %s", err)
	}
}
