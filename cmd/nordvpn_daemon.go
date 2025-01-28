package cmd

import (
	"fmt"
	"log/slog"
	"os/exec"
)

// NordVPNDaemon --
type NordVPNDaemon struct{ next IBootUP }

func (n *NordVPNDaemon) setNext(next IBootUP) { n.next = next }

func (n *NordVPNDaemon) execute(p *BootUPParams) {

	// set inital daemon state TRUE
	p.IsDaemonRunning = true

	// 1. Check if nordvpn daemon is running
	slog.Info("Checking nordvpn daemon...")
	out, err := nordVPNDaemonStatus.Output()
	slog.Info(string(out))
	if err != nil {
		p.IsDaemonRunning = false
	}

	if !p.IsDaemonRunning {

		// 1.1 If nordvpn not running start daemon
		out, err = nordVPNDaemonStart.Output()
		slog.Info(fmt.Sprintf("Starting nordvpn daemon.. %s", string(out)))
		if err != nil {
			panic(err)
		}

		// 1.2 Wait 5 second until daemon started
		_, err = exec.Command("sleep", "5").Output()
		if err != nil {
			panic(err)
		}

		// 1.3 Maksure daemon running
		out, err = nordVPNDaemonStatus.Output()
		slog.Info(string(out))
		if err != nil {
			panic(err)
		}

		p.IsDaemonRunning = true
	}
	slog.Info("Nordvpn daemon started")

	if n.next != nil {
		n.next.execute(p)
	}
}
