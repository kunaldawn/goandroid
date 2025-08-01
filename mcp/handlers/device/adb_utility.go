package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AdbUtilityHandler handles ADB utility operations
type AdbUtilityHandler struct {
	manager goandroid.AndroidManager
}

// NewAdbUtilityHandler creates a new ADB utility handler
func NewAdbUtilityHandler(manager goandroid.AndroidManager) *AdbUtilityHandler {
	return &AdbUtilityHandler{manager: manager}
}

// RegisterTools registers ADB utility operation tools
func (h *AdbUtilityHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Wait for specific serials
	waitForSerialsTool := mcpLib.NewTool("wait_for_serials",
		mcpLib.WithDescription("Wait for specific device serials to become available"),
		mcpLib.WithNumber("timeout",
			mcpLib.Required(),
			mcpLib.Description("Timeout in seconds"),
		),
		mcpLib.WithArray("serials",
			mcpLib.Required(),
			mcpLib.Description("Array of device serial numbers to wait for"),
		),
	)
	mcpServer.AddTool(waitForSerialsTool, h.handleWaitForSerials)

	// Wait for device count
	waitForDeviceCountTool := mcpLib.NewTool("wait_for_device_count",
		mcpLib.WithDescription("Wait for a specific number of devices to be available"),
		mcpLib.WithNumber("timeout",
			mcpLib.Required(),
			mcpLib.Description("Timeout in seconds"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Required(),
			mcpLib.Description("Number of devices to wait for"),
		),
	)
	mcpServer.AddTool(waitForDeviceCountTool, h.handleWaitForDeviceCount)
}

func (h *AdbUtilityHandler) handleWaitForSerials(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Timeout float64  `json:"timeout"`
		Serials []string `json:"serials"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	timeout := int(args.Timeout)
	if timeout <= 0 {
		return mcpLib.NewToolResultError("Timeout must be positive"), nil
	}

	if len(args.Serials) == 0 {
		return mcpLib.NewToolResultError("At least one serial must be provided"), nil
	}

	err := h.manager.Endpoint.WaitForSerials(timeout, args.Serials...)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "All specified serials are now available",
		"serials": args.Serials,
		"timeout": timeout,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdbUtilityHandler) handleWaitForDeviceCount(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Timeout float64 `json:"timeout"`
		Count   float64 `json:"count"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	timeout := int(args.Timeout)
	count := int(args.Count)

	if timeout <= 0 {
		return mcpLib.NewToolResultError("Timeout must be positive"), nil
	}

	if count <= 0 {
		return mcpLib.NewToolResultError("Count must be positive"), nil
	}

	err := h.manager.Endpoint.WaitForDevices(timeout, count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("%d devices are now available", count),
		"count":   count,
		"timeout": timeout,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}