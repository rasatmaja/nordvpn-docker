package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define signal interrupt
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// define error group
	proc, gctx := errgroup.WithContext(ctx)

	// wait for signal
	proc.Go(func() error {
		sig := <-sigs

		log.Printf("received signal %s \n", sig)

		// cancel context after reveived signal interrupt
		cancel()

		return nil
	})

	proc.Go(func() error {
		// start nordvpn daemon
		log.Println("Starting nordvpn daemon...")
		_, err := exec.CommandContext(gctx, "/etc/init.d/nordvpn", "start").Output()
		if err != nil {
			return err
		}

		return nil
	})

	if err := proc.Wait(); err != nil {
		log.Panicln(err)
	}

	log.Println("Shutting down...")
	log.Println("Shutdown complete")
}
