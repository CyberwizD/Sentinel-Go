# Sentinel-Go

⚡ A lightweight and fast **Web API Security Scanner** built in **Go**. It helps you quickly identify common API security misconfigurations and weaknesses such as missing HTTPS, missing security headers, server information leaks, and basic injection vulnerabilities.

---

## Main Idea:
Command-line tool in Go that accepts a URL (an API endpoint) and checks for basic security issues like:

* Missing HTTPS

* HTTP headers missing (like Content-Security-Policy, Strict-Transport-Security, X-Content-Type-Options, etc.)

* Insecure Server Banner leaks (like “Apache 2.2.3”)

* Weak Authentication (e.g., WWW-Authenticate missing)

* Basic SQLi test by injecting payloads into query parameters (like id=1' and see the response)

* JSON Response Content-Type Check (application/json expected)

🔗 (It's like your tiny custom BurpSuite!)

---

## 🛠 Why it’s PERFECT:

✅ Uses Golang HTTP client (good practice)

✅ Practice secure programming (timeouts, retries, concurrency)

✅ Learn real-world security checks

✅ Portfolio shows you can build security tools, not just “learn theory”

✅ Can easily extend later (to add more pentest modules)

--- 

## ✨ Features

- 🔒 HTTPS Enforcement Check
- 🛡 Security Headers Validation
- 🕵️ Server Banner Information Leak Detection
- 💥 Basic SQL Injection Detection (coming soon)
- 📄 Save Scan Reports to JSON (optional)

---

## 📦 Project Structure

```
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
```

---

## 🚀 Quick Start

1. Clone the repository:

```bash
git clone https://github.com/yourusername/Sentinel-Go.git
cd Sentinel-Go
```

2. Install dependencies and run:

```go
go run root.go -url https://example.com/api/v1/users
```

## 🛠 How It Works
* Connects to the target URL

* Checks HTTPS enforcement

* Analyzes important HTTP security headers

* Detects server information disclosure

* (Optionally) injects basic SQL payloads into parameters

## 📜 Example Usage:

```go
$ Sentinel-Go -url https://target.com/api/user?id=1

[+] Checking HTTPS... OK
[+] Checking Security Headers... Missing: Content-Security-Policy
[+] Checking Server Info Leak... Detected: "Apache/2.4.41"
[+] Testing for SQL Injection... Possible anomaly detected!
[+] Validating Content-Type... OK (application/json)
```

---

## 🧩 Future Improvements
* Threaded scanning for multiple endpoints

* SQLi and XSS payload testing

* Beautiful CLI output with color

* Web UI for visual scan reports

---

## 🤝 Contribution
Contributions, issues, and feature requests are welcome!
Feel free to fork and submit pull requests.
