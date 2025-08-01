package display

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DisplayHandler handles display management operations
type DisplayHandler struct {
	manager goandroid.AndroidManager
}

// NewDisplayHandler creates a new display handler
func NewDisplayHandler(manager goandroid.AndroidManager) *DisplayHandler {
	return &DisplayHandler{manager: manager}
}

// RegisterTools registers display management tools
func (h *DisplayHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get display size tool
	getDisplaySizeTool := mcpLib.NewTool("get_display_size",
		mcpLib.WithDescription("Get the current display size of the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getDisplaySizeTool, h.handleGetDisplaySize)

	// Set display size tool
	setDisplaySizeTool := mcpLib.NewTool("set_display_size",
		mcpLib.WithDescription("Set the display size of the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("width",
			mcpLib.Required(),
			mcpLib.Description("Width in pixels"),
		),
		mcpLib.WithNumber("height",
			mcpLib.Required(),
			mcpLib.Description("Height in pixels"),
		),
	)
	mcpServer.AddTool(setDisplaySizeTool, h.handleSetDisplaySize)
}

// handleGetDisplaySize handles the get_display_size tool
func (h *DisplayHandler) handleGetDisplaySize(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	width, height, err := device.Display.GetDisplaySize()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get display size: %v", err)), nil
	}

	result := map[string]interface{}{
		"width":  width,
		"height": height,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleSetDisplaySize handles the set_display_size tool
func (h *DisplayHandler) handleSetDisplaySize(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	widthVal, ok := args["width"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("width must be a number"), nil
	}
	width := int(widthVal)

	heightVal, ok := args["height"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("height must be a number"), nil
	}
	height := int(heightVal)

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Display.SetDisplaySize(width, height)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to set display size: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Display size set to %dx%d", width, height),
		"width":   width,
		"height":  height,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}