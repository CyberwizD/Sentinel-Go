package scanner

import (
	"fmt"
	"net/http"

	"github.com/CyberwizD/Sentinel-Go/utils"
)

func RunBasicScans(url string) {
	client := utils.NewHTTPClient()

	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("[-] Failed to connect: %v\n", err)
		return
	}

	defer resp.Body.Close()

	checkHTTPS(url)
	checkSecurityHeaders(resp)
	checkServerLeaks(resp)
}

func checkHTTPS(url string) {
	if url[:5] != "https" {
		fmt.Println("[-] Warning: URL is not using HTTPS.")
	} else {
		fmt.Println("[+] HTTPS check passed.")
	}
}

func checkSecurityHeaders(resp *http.Response) {
	requiredHeaders := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
	}

	fmt.Println("[*] Checking security headers...")

	for _, header := range requiredHeaders {
		if resp.Header.Get(header) == "" {
			fmt.Printf("[-] Missing header: %s\n", header)
		} else {
			fmt.Printf("[+] Header present: %s\n", header)
		}
	}
}

func checkServerLeaks(resp *http.Response) {
	server := resp.Header.Get("Server")

	if server != "" {
		fmt.Printf("[-] Server banner detected: %s\n", server)
	} else {
		fmt.Println("[+] No server banner detected.")
	}
}
