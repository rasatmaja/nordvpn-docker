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

	// 1. Check if nordvpn daemon is running
	log.Println("Checking nordvpn daemon...")
	out, err := exec.Command("/etc/init.d/nordvpn", "status").Output()
	if err != nil {
		log.Printf("%s", out)

		// 1.1 If nordvpn not running start daemon
		log.Println("Starting nordvpn daemon...")
		_, err := exec.Command("/etc/init.d/nordvpn", "start").Output()
		if err != nil {
			log.Panic(err)
		}
	}
	log.Println("Nordvpn daemon started")

	// 2. Check if account already loggen in
	log.Println("Checking nordvpn account...")
	out, err = exec.Command("nordvpn", "account").Output()
	if err != nil {
		// 2.1 If account are no logged in then try login using token
		log.Printf("%s", out)
		out, err = exec.Command("nordvpn", "login", "--token", os.Getenv("NORDVPN_TOKEN")).Output()
		if err != nil {
			log.Printf("%s", out)
			log.Panic(err)
		}
		log.Printf("%s", out)
	}

	log.Println("Checking nordvpn status...")
	out, err = exec.Command("nordvpn", "status").Output()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("%s", out)

	log.Printf("Connecting nordvpn to %s \n", os.Getenv("NORDVPN_DEFAULT_CONNECT_COUNTRY"))
	out, err = exec.Command("nordvpn", "c", os.Getenv("NORDVPN_DEFAULT_CONNECT_COUNTRY")).Output()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%s", out)

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
