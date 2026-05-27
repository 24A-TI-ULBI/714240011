package helper

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ResponseFormat defines the standard API response structure
type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(ResponseFormat{
		Status: "success",
		Data:   data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ResponseFormat{
		Status:  "error",
		Message: message,
	})
}

// GetAddress reads PORT and IP from env, returns listen address and network type.
// Mengikuti pola boilerplate gocroot — support IPv4 dan IPv6.
func GetAddress() (ipport string, network string) {
	port := os.Getenv("PORT")
	network = "tcp4"
	if port == "" {
		ipport = ":8080"
	} else if port[0:1] != ":" {
		ip := os.Getenv("IP")
		if ip == "" {
			ipport = ":" + port
		} else {
			if strings.Contains(ip, ".") {
				ipport = ip + ":" + port
			} else {
				ipport = "[" + ip + "]" + ":" + port
				network = "tcp6"
			}
		}
	}
	return
}

// GetIPaddress fetches the public IP address of the server via icanhazip.com
func GetIPaddress() string {
	resp, err := http.Get("https://icanhazip.com/")
	if err != nil {
		log.Println("Gagal mengambil IP publik:", err)
		return "unknown"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Gagal membaca response IP:", err)
		return "unknown"
	}
	return strings.TrimSpace(string(body))
}
