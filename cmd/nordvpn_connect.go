package cmd

import (
	"log/slog"
)

// NordVPNConnect --
type NordVPNConnect struct{ next IBootUP }

func (n *NordVPNConnect) setNext(next IBootUP) { n.next = next }

func (n *NordVPNConnect) execute(p *BootUPParams) {
	out, err := nordVPNAppConnect.Output()
	slog.Info(string(out))
	if err != nil {
		panic(err)
	}

	if n.next != nil {
		n.next.execute(p)
	}
}
