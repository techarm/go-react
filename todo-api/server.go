package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv      *http.Server
	listener net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv:      &http.Server{Handler: mux},
		listener: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	// 別のゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
		}
		return nil
	})

	// チャネルからの終了通知を待機する
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// 別のゴルーチンの終了をまつ
	return eg.Wait()
}
