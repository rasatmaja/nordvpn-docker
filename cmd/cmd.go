package cmd

import (
	"os"
	"os/exec"
)

var (
	// NordVPN Daemon
	nordVPNDaemonStatus = command("/etc/init.d/nordvpn", "status")
	nordVPNDaemonStart  = command("/etc/init.d/nordvpn", "start")

	// NordVPN App
	nordVPNAppStatus  = command("nordvpn", "status")
	nordVPNAppAccount = command("nordvpn", "account")
	nordVPNAppLogin   = command("nordvpn", "login", "--token", os.Getenv("NORDVPN_TOKEN"))
	nordVPNAppConnect = command("nordvpn", "c", os.Getenv("NORDVPN_DEFAULT_CONNECT_COUNTRY"))

	// NordVPN Config
	nordVPNAppEnableLANDiscovery = command("nordvpn", "set", "lan-discovery", os.Getenv("NORDVPN_ENABLE_LAN_DISCOVERY"))
)

type cmd struct {
	name string
	args []string
}

func command(name string, args ...string) cmd {
	return cmd{name: name, args: args}
}

func (c cmd) Output() ([]byte, error) {
	return exec.Command(c.name, c.args...).Output()
}
