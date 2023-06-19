package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/edge/api"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	configurations, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	flag.Parse()

	cmd := flag.Arg(0)

	go func() {
		switch cmd {
		case "serve":
			api.StartServer(ctx, configurations)
			cancel()
		default:
			panic(errors.New("wrong command"))
		}

	}()

	sig := make(chan os.Signal, 10)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sig:
		fmt.Printf("Received %s, graceful shut down...", s.String())
		cancel()
	case <-ctx.Done():
	}

	time.Sleep(1 * time.Second)
}
