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

	// NordVPN Settings
	nordVPNAppEnableLANDiscovery = command("nordvpn", "set", "lan-discovery", os.Getenv("NORDVPN_ENABLE_LAN_DISCOVERY"))
	nordVPNAppEnableKillSwitch   = command("nordvpn", "set", "killswitch", os.Getenv("NORDVPN_ENABLE_KILL_SWITCH"))
	nordVPNAppEnableIPv6         = command("nordvpn", "set", "ipv6", os.Getenv("NORDVPN_ENABLE_IPV6"))
	nordVPNAppEnableFirewall     = command("nordvpn", "set", "firewall", os.Getenv("NORDVPN_ENABLE_FIREWALL"))
	nordVPNAppTechnology         = command("nordvpn", "set", "technology", os.Getenv("NORDVPN_DEFAULT_TECHNOLOGY"))
	nordVPNAppEnableAutoConnect  = command("nordvpn", "set", "autoconnect", os.Getenv("NORDVPN_ENABLE_AUTO_CONNECT"))
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
