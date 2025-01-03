package cmd

import (
	"log/slog"
)

// NordVPNAccount --
type NordVPNAccount struct{ next IBootUP }

func (n *NordVPNAccount) setNext(next IBootUP) { n.next = next }

func (n *NordVPNAccount) execute(p *BootUPParams) {

	p.IsAccountLoggedIn = true

	// 2. Check if account already loggen in
	slog.Info("Checking nordvpn account...")
	out, err := nordVPNAppAccount.Output()
	slog.Info(string(out))
	if err != nil {
		p.IsAccountLoggedIn = false
	}

	if !p.IsAccountLoggedIn {
		// 2.1 If account are no logged in then try login using token
		slog.Info(string(out))
		out, err = nordVPNAppLogin.Output()
		if err != nil {
			slog.Error(string(out))
			panic(err)
		}
		slog.Info(string(out))
		p.IsAccountLoggedIn = true
	}

	if n.next != nil {
		n.next.execute(p)
	}
}
