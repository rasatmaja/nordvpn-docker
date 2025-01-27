package main

import (
	"context"
	"log/slog"
	"nordvpn-docker/cmd"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	// get args without program name
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "healthcheck" {
		healthcheck()
		return
	}

	bootup()
}

func healthcheck() {
	cmd.Healtcheck()
}

func bootup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.BootUP()

	// define signal interrupt
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// define error group
	proc, _ := errgroup.WithContext(ctx)

	// wait for signal
	proc.Go(func() error {
		sig := <-sigs

		slog.Info("received signal", slog.String("signal", sig.String()))

		// cancel context after reveived signal interrupt
		cancel()

		return nil
	})

	if err := proc.Wait(); err != nil {
		panic(err)
	}

	slog.Info("Shutting down...")
	slog.Info("Shutdown complete")
}
