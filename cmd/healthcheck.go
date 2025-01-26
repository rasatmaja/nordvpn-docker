package cmd

import (
	"log/slog"
	"os"
	"strings"
)

// Healtcheck --
func Healtcheck() {
	out, err := nordVPNAppStatus.Output()
	if err != nil {
		slog.Error(string(out))
		os.Exit(1)
	}

	data := parseStatus(string(out))
	status := data["status"]

	if status != "Connected" {
		slog.Error("NordVPN Not Connected", slog.String("status", status))
		os.Exit(1)
	}

	server := data["server"]
	ip := data["ip"]
	slog.Info("NordVPN Connected", slog.String("server", server), slog.String("ip", ip))
}

// Parse output of NordVPN status
// ========Example raw output:==========
// Status: Connected
// Server: Singapore #123
// Hostname: xxxx.nordvpn.com
// IP: 192.166.246.xxx
// Country: Singapore
// City: Singapore
// Current technology: NORDLYNX
// Current protocol: UDP
// Post-quantum VPN: Disabled
// Transfer: 92 B received, 180 B sent
// Uptime: 1 second
// =====================================
func parseStatus(input string) map[string]string {
	lines := strings.Split(input, "\n")
	status := make(map[string]string, len(lines))

	for _, line := range lines {
		data := strings.Split(line, ":")
		if len(data) > 1 {
			key := strings.ToLower(data[0])
			value := strings.TrimSpace(data[1])
			status[key] = value
		}
	}
	return status
}
