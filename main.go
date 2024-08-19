package main

import (
	"context"
	"log"
	"nordvpn-docker/cmd"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

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

		log.Printf("received signal %s \n", sig)

		// cancel context after reveived signal interrupt
		cancel()

		return nil
	})

	if err := proc.Wait(); err != nil {
		log.Panicln(err)
	}

	log.Println("Shutting down...")
	log.Println("Shutdown complete")
}
