package main

import (
	"flag"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CyberwizD/Sentinel-Go/scanner"
)

func Run() {
	url := flag.String("url", "", "Target URL to scan")
	flag.Parse()

	if *url == "" {
		fmt.Println("[-] Please provide a URL using the -url flag.")
		return
	}

	scanner.RunBasicScans(*url)

	cobraCmd := &cobra.Command{
		Use:   "sentinel-go",
		Short: "Sentinel-Go: A simple web API request security scanner",
		Long:  `Sentinel-Go is a simple web API security scanner that checks for common vulnerabilities and security issues.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Sentinel-Go is running...")
		},
	}

	cobraCmd.Flags().StringVarP(url, "url", "u", "", "Target URL to scan")
	cobraCmd.MarkFlagRequired("url")
}

func main() {
	Run()
}
