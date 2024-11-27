package cmd

// BootUP --
func BootUP() {
	var params *BootUPParams
	params = new(BootUPParams)

	var daemon, account, connect, config IBootUP
	daemon = new(NordVPNDaemon)
	account = new(NordVPNAccount)
	connect = new(NordVPNConnect)
	config = new(NordVPNConfig)

	// set config flow
	connect.setNext(config)

	// set bootup flow
	account.setNext(connect)
	daemon.setNext(account)

	// execute
	daemon.execute(params)
}

// IBootUP --
type IBootUP interface {
	execute(*BootUPParams)
	setNext(IBootUP)
}

// BootUPParams --
type BootUPParams struct {
	IsDaemonRunning   bool
	IsAccountLoggedIn bool
}
