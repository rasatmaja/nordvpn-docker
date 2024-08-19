package cmd

import (
	"log"
)

// NordVPNAccount --
type NordVPNAccount struct{ next IBootUP }

func (n *NordVPNAccount) setNext(next IBootUP) { n.next = next }

func (n *NordVPNAccount) execute(p *BootUPParams) {

	p.IsAccountLoggedIn = true

	// 2. Check if account already loggen in
	log.Println("Checking nordvpn account...")
	out, err := nordVPNAppAccount.Output()
	log.Printf("%s", out)
	if err != nil {
		p.IsAccountLoggedIn = false
	}

	if !p.IsAccountLoggedIn {
		// 2.1 If account are no logged in then try login using token
		log.Printf("%s", out)
		out, err = nordVPNAppLogin.Output()
		if err != nil {
			log.Printf("%s", out)
			log.Panic(err)
		}
		log.Printf("%s", out)
		p.IsAccountLoggedIn = true
	}

	if n.next != nil {
		n.next.execute(p)
	}
}
