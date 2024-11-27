package cmd

import (
	"log"
)

// NordVPNConfig --
type NordVPNConfig struct{ next IBootUP }

func (n *NordVPNConfig) setNext(next IBootUP) { n.next = next }

func (n *NordVPNConfig) execute(p *BootUPParams) {

	for cfg, funcx := range configs {
		log.Printf("set config %s", cfg)
		funcx(p)
	}

	if n.next != nil {
		n.next.execute(p)
	}
}

var configs = map[string]func(p *BootUPParams){
	"lan-discovery": func(p *BootUPParams) {
		if !p.IsDaemonRunning {
			log.Panic("Daemon not running")
		}

		// set nordvpn lan discovery config
		out, err := nordVPNAppEnableLANDiscovery.Output()
		log.Printf("%s", out)
		if err != nil {
			log.Panic(err)
		}
	},
}
