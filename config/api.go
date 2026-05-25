package config

import (
	"bufio"
	"os"
	"strings"
)

var _ = loadEnv(".env")

func loadEnv(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		// Remove quotes
		if (strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"")) ||
			(strings.HasPrefix(val, "'") && strings.HasSuffix(val, "'")) {
			val = val[1 : len(val)-1]
		}
		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}
	return true
}

var WAKeyword string = os.Getenv("WAQRKEYWORD")

var WebhookURL string = os.Getenv("WEBHOOKURL")

var WebhookSecret string = os.Getenv("WEBHOOKSECRET")

var WAPhoneNumber string = os.Getenv("WAPHONENUMBER")

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/send/message/text"

var WAAPIGetToken string = "https://api.wa.my.id/api/signup"
