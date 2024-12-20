package cmd

// BootUP --
func BootUP() {
	var params *BootUPParams
	params = new(BootUPParams)

	var daemon, account, connect, setting IBootUP
	daemon = new(NordVPNDaemon)
	account = new(NordVPNAccount)
	connect = new(NordVPNConnect)
	setting = new(NordVPNSetting)

	// set setting flow
	connect.setNext(setting)

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
