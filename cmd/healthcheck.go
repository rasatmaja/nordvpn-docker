package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

// Healtcheck --
func Healtcheck() {
	out, err := nordVPNAppStatus.Output()
	if err != nil {
		slog.Error(string(out))
		os.Exit(1)
	}
	status := findStatus("Status", string(out))
	server := findStatus("Server", string(out))
	ip := findStatus("IP", string(out))

	if status != "Connected" {
		slog.Error("NordVPN Not Connected", slog.String("status", status))
		os.Exit(1)
	}

	slog.Info("NordVPN Connected", slog.String("server", server), slog.String("ip", ip))
}

const pattern = `(?i)%s:\s*(.*)`

func findStatus(field, input string) string {
	pattern := fmt.Sprintf(pattern, field)
	key := regexp.MustCompile(pattern)
	output := key.FindStringSubmatch(input)

	if len(output) > 0 {
		return output[1]
	}
	return ""
}
