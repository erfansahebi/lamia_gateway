package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/edge/api"
	sharedCommon "github.com/erfansahebi/lamia_shared/common"
	"github.com/erfansahebi/lamia_shared/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	configurations, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatalf(ctx, "failed to load config")
		panic(err)
	}

	flag.Parse()

	cmd := flag.Arg(0)

	log.Infof(ctx, "Starting with command: %s", cmd)

	go func() {
		switch cmd {
		case "serve":
			api.StartServer(ctx, configurations)
			cancel()
		default:
			log.WithError(sharedCommon.ErrWrongCommand).Fatalf(ctx, "wrong command")
			panic(sharedCommon.ErrWrongCommand)
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
