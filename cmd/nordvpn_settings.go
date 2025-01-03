package cmd

import (
	"log/slog"
)

// NordVPNSetting --
type NordVPNSetting struct{ next IBootUP }

func (n *NordVPNSetting) setNext(next IBootUP) { n.next = next }

func (n *NordVPNSetting) execute(p *BootUPParams) {

	for setting, funcx := range settings {
		slog.Info("set setting", slog.String("setting", setting))
		funcx(p)
	}

	if n.next != nil {
		n.next.execute(p)
	}
}

var settings = map[string]func(p *BootUPParams){
	"lan-discovery": func(p *BootUPParams) {
		if !p.IsDaemonRunning {
			panic("Daemon not running")
		}

		// set nordvpn lan discovery settings
		out, err := nordVPNAppEnableLANDiscovery.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
	"killswitch": func(p *BootUPParams) {
		// set nordvpn kill switch settings
		out, err := nordVPNAppEnableKillSwitch.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
	"ipv6": func(p *BootUPParams) {
		out, err := nordVPNAppEnableIPv6.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
	"firewall": func(p *BootUPParams) {
		out, err := nordVPNAppEnableFirewall.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
	"technology": func(p *BootUPParams) {
		out, err := nordVPNAppTechnology.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
	"autoconnect": func(p *BootUPParams) {
		out, err := nordVPNAppEnableAutoConnect.Output()
		slog.Info(string(out))
		if err != nil {
			slog.Error(err.Error())
		}
	},
}
