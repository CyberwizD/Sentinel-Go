# Sentinel-Go

⚡ A lightweight and fast **Web API Security Scanner** built in **Go**. It helps you quickly identify common API security misconfigurations and weaknesses such as missing HTTPS, missing security headers, server information leaks, and basic injection vulnerabilities.

---

## ✨ Features

- 🔒 HTTPS Enforcement Check
- 🛡 Security Headers Validation
- 🕵️ Server Banner Information Leak Detection
- 💥 Basic SQL Injection Detection (coming soon)
- 📄 Save Scan Reports to JSON (optional)

---

## 📦 Project Structure

Sentinel-Go/
├── cmd/
│   └── root.go      # CLI flags, main app entry
├── internal/
│   ├── scanner/
│   │   ├── scanner.go   # Core scanning functions (HTTPS check, headers check, etc.)
│   │   └── payloads.go  # Payloads for SQLi and other tests
│   └── utils/
│       └── httpclient.go # Custom HTTP client (timeouts, retries, etc.)
├── reports/
│   └── (outputs JSON reports)
├── go.mod
├── go.sum
└── README.md

---

## 🚀 Quick Start

1. Clone the repository:

```bash
git clone https://github.com/yourusername/Sentinel-Go.git
cd Sentinel-Go
```

2. Install dependencies and run:

```go
go run main.go -url https://example.com/api/v1/users
```

## 🛠 How It Works
* Connects to the target URL

* Checks HTTPS enforcement

* Analyzes important HTTP security headers

* Detects server information disclosure

* (Optionally) injects basic SQL payloads into parameters

## 📜 Example Usage:

```go
$ go run main.go -url https://api.example.com/login

[+] HTTPS check passed.
[*] Checking security headers...
[-] Missing header: Content-Security-Policy
[-] Missing header: X-Content-Type-Options
[-] Server banner detected: Apache/2.4.41
```

## 🧩 Future Improvements
* Threaded scanning for multiple endpoints

* SQLi and XSS payload testing

* Beautiful CLI output with color

* Web UI for visual scan reports

## 🤝 Contribution
Contributions, issues, and feature requests are welcome!
Feel free to fork and submit pull requests.
