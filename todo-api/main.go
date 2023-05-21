package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/techarm/todo-api/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("failed to teminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 設定情報を読み込む
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {
		return err
	}

	s := NewServer(l, mux)
	return s.Run(ctx)
}
