package cmd

import (
	"log"
)

// NordVPNSetting --
type NordVPNSetting struct{ next IBootUP }

func (n *NordVPNSetting) setNext(next IBootUP) { n.next = next }

func (n *NordVPNSetting) execute(p *BootUPParams) {

	for setting, funcx := range settings {
		log.Printf("set settings %s", setting)
		funcx(p)
	}

	if n.next != nil {
		n.next.execute(p)
	}
}

var settings = map[string]func(p *BootUPParams){
	"lan-discovery": func(p *BootUPParams) {
		if !p.IsDaemonRunning {
			log.Panic("Daemon not running")
		}

		// set nordvpn lan discovery settings
		out, err := nordVPNAppEnableLANDiscovery.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Println(err)
		}
	},
	"killswitch": func(p *BootUPParams) {
		// set nordvpn kill switch settings
		out, err := nordVPNAppEnableKillSwitch.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Println(err)
		}
	},
	"ipv6": func(p *BootUPParams) {
		out, err := nordVPNAppEnableIPv6.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Println(err)
		}
	},
	"firewall": func(p *BootUPParams) {
		out, err := nordVPNAppEnableFirewall.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Println(err)
		}
	},
}
