package Week03

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	svc := NewService(os.Stderr)

	if err := svc.Start(); err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case sig := <-signalChan:
			log.Printf("Graceful shutdown service due to signal: %s", sig.String())
			if err := svc.Stop(); err != nil {
				log.Fatalf("Shutdown service with error: %v", err)
			}

			break loop
		case err := <-svc.ErrChan():
			log.Fatalf("Service stopped due to error: %v", err)
		}
	}

	log.Println("service is stopped")
}

type Service struct {
	w           io.Writer
	httpServers []*http.Server
	errChan     chan error
}

func NewService(w io.Writer) *Service {
	return &Service{
		w:       w,
		errChan: make(chan error, 1),
	}
}

func (s *Service) Start() error {
	public := &http.Server{
		Addr:              "localhost:8080",
		Handler:           nil,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	admin := &http.Server{
		Addr:              "localhost:8081",
		Handler:           nil,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	s.httpServers = []*http.Server{public, admin}
	s.errChan <- public.ListenAndServe()
	s.errChan <- admin.ListenAndServe()

	select {
	case err := <-s.errChan:
		return err
	default:
		return nil
	}
}

func (s *Service) Stop() error {
	g, ctx := errgroup.WithContext(context.Background())

	for _, svc := range s.httpServers {
		svc := svc
		g.Go(func() error {
			return svc.Shutdown(ctx)
		})
	}

	return g.Wait()
}

func (s *Service) ErrChan() <-chan error {
	return s.errChan
}
