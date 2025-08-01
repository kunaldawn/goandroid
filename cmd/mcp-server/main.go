package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kunaldawn/goandroid/mcp"
)

func main() {
	var (
		timeout  = flag.Int("timeout", 30, "ADB operation timeout in seconds")
		adbPath  = flag.String("adb", "adb", "Path to ADB executable")
		help     = flag.Bool("help", false, "Show help message")
		httpAddr = flag.String("http", "", "HTTP server address (e.g., :8080). If not set, uses stdio")
	)
	flag.Parse()

	if *help {
		fmt.Printf("GoAndroid MCP Server\n\n")
		fmt.Printf("This server exposes goandroid automation capabilities via the Model Context Protocol (MCP).\n")
		fmt.Printf("It provides tools for Android device automation including input, UI interaction, and activity management.\n\n")
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		fmt.Printf("Example:\n")
		fmt.Printf("  %s --timeout 60 --adb /path/to/adb           # Stdio mode\n", os.Args[0])
		fmt.Printf("  %s --http :8080                                # HTTP mode on port 8080\n\n", os.Args[0])
		return
	}

	log.Printf("Starting GoAndroid MCP Server...")
	log.Printf("ADB timeout: %d seconds", *timeout)
	log.Printf("ADB path: %s", *adbPath)

	server := mcp.NewGoAndroidMCPServer(*timeout, *adbPath)

	if *httpAddr != "" {
		log.Printf("Server initialized. Starting HTTP transport on %s...", *httpAddr)
		if err := server.ServeHTTP(*httpAddr); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		log.Printf("Server initialized. Starting stdio transport...")
		if err := server.ServeStdio(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}
}
