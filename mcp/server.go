package mcp

import (
	"github.com/kunaldawn/goandroid"
	"github.com/kunaldawn/goandroid/mcp/registry"
	"github.com/mark3labs/mcp-go/server"
)

// NewGoAndroidMCPServer creates a new MCP server instance with goandroid integration
func NewGoAndroidMCPServer(timeout int, adbPath string) *GoAndroidMCPServer {
	s := server.NewMCPServer(
		"goandroid-mcp-server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	manager := goandroid.GetNewAndroidManager(timeout, adbPath)

	mcpServer := &GoAndroidMCPServer{
		server:  s,
		manager: manager,
	}

	// Register all tools using the centralized registry
	handlerRegistry := registry.NewHandlerRegistry(manager)
	handlerRegistry.RegisterAllTools(s)

	return mcpServer
}

// ServeStdio starts the MCP server using stdio transport
func (s *GoAndroidMCPServer) ServeStdio() error {
	return server.ServeStdio(s.server)
}

// ServeHTTP starts the MCP server using HTTP transport with streaming support
func (s *GoAndroidMCPServer) ServeHTTP(addr string) error {
	httpServer := server.NewStreamableHTTPServer(s.server)
	return httpServer.Start(addr)
}
