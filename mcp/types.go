package mcp

import (
	"github.com/kunaldawn/goandroid"
	"github.com/mark3labs/mcp-go/server"
)

// GoAndroidMCPServer wraps the MCP server with goandroid functionality
type GoAndroidMCPServer struct {
	server  *server.MCPServer
	manager goandroid.AndroidManager
}

// ToolHandler defines the interface for MCP tool handlers
type ToolHandler interface {
	RegisterTools(server *GoAndroidMCPServer)
}