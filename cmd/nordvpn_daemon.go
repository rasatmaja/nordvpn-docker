package cmd

import (
	"log"
	"os/exec"
)

// NordVPNDaemon --
type NordVPNDaemon struct{ next IBootUP }

func (n *NordVPNDaemon) setNext(next IBootUP) { n.next = next }

func (n *NordVPNDaemon) execute(p *BootUPParams) {

	// set inital daemon state TRUE
	p.IsDaemonRunning = true

	// 1. Check if nordvpn daemon is running
	log.Println("Checking nordvpn daemon...")
	out, err := nordVPNDaemonStatus.Output()
	log.Printf("%s", out)
	if err != nil {
		p.IsDaemonRunning = false
	}

	if !p.IsDaemonRunning {

		// 1.1 If nordvpn not running start daemon
		log.Println("Starting nordvpn daemon...")
		_, err = nordVPNDaemonStart.Output()
		if err != nil {
			log.Panic(err)
		}

		// 1.2 Wait 5 second until daemon started
		_, err = exec.Command("sleep", "5").Output()
		if err != nil {
			log.Panic(err)
		}

		// 1.3 Maksure daemon running
		out, err = nordVPNDaemonStatus.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Panic(err)
		}

		p.IsDaemonRunning = true
	}
	log.Println("Nordvpn daemon started")

	if n.next != nil {
		n.next.execute(p)
	}
}
