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

type server struct {
	db dber
}

func newServer() *server {
	return &server{
		db: newDB(),
	}
}

func main() {
	addr := flag.String("addr", ":8080", "TCP address to listen to")
	flag.Parse()

	s := newServer()
	r := routes(s)

	fs := &fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		log.Printf("starting server at %q", *addr)
		if err := fs.ListenAndServe(*addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	log.Println("shutting down server")
	if err := fs.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("shutting down server: %s", err)
	}
}
