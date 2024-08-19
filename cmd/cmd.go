package cmd

import (
	"os"
	"os/exec"
)

var (
	// NordVPN Daemon
	nordVPNDaemonStatus = exec.Command("/etc/init.d/nordvpn", "status")
	nordVPNDaemonStart  = exec.Command("/etc/init.d/nordvpn", "start")

	// NordVPN App
	nordVPNAppStatus  = exec.Command("nordvpn", "status")
	nordVPNAppAccount = exec.Command("nordvpn", "account")
	nordVPNAppLogin   = exec.Command("nordvpn", "login", "--token", os.Getenv("NORDVPN_TOKEN"))
	nordVPNAppConnect = exec.Command("nordvpn", "c", os.Getenv("NORDVPN_DEFAULT_CONNECT_COUNTRY"))
)
