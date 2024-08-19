package cmd

import (
	"log"
)

// NordVPNConnect --
type NordVPNConnect struct{ next IBootUP }

func (n *NordVPNConnect) setNext(next IBootUP) { n.next = next }

func (n *NordVPNConnect) execute(p *BootUPParams) {
	out, err := nordVPNAppConnect.Output()
	log.Printf("%s", out)
	if err != nil {
		log.Panic(err)
	}

	if n.next != nil {
		n.next.execute(p)
	}
}
