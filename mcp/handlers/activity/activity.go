package activity

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ActivityHandler handles activity management operations
type ActivityHandler struct {
	manager goandroid.AndroidManager
}

// NewActivityHandler creates a new activity handler
func NewActivityHandler(manager goandroid.AndroidManager) *ActivityHandler {
	return &ActivityHandler{manager: manager}
}

// RegisterTools registers activity management tools
func (h *ActivityHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Launch activity tool
	launchActivityTool := mcpLib.NewTool("launch_activity",
		mcpLib.WithDescription("Launch an Android activity"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("activity",
			mcpLib.Required(),
			mcpLib.Description("Canonical class name of the activity to launch"),
		),
	)
	mcpServer.AddTool(launchActivityTool, h.handleLaunchActivity)

	// Get focused activity tool
	getFocusedActivityTool := mcpLib.NewTool("get_focused_activity",
		mcpLib.WithDescription("Get the currently focused activity"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getFocusedActivityTool, h.handleGetFocusedActivity)

	// Wait for activity tool
	waitForActivityTool := mcpLib.NewTool("wait_for_activity",
		mcpLib.WithDescription("Wait for a specific activity to become focused"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("activity",
			mcpLib.Required(),
			mcpLib.Description("Activity name to wait for"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(waitForActivityTool, h.handleWaitForActivity)
}

// handleLaunchActivity handles the launch_activity tool
func (h *ActivityHandler) handleLaunchActivity(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	activity, err := request.RequireString("activity")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Activity.StartActivity(activity)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to launch activity %s: %v", activity, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Launched activity: %s", activity),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleGetFocusedActivity handles the get_focused_activity tool
func (h *ActivityHandler) handleGetFocusedActivity(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	focusedActivity, err := device.Activity.GetFocusedActivity()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get focused activity: %v", err)), nil
	}

	result := map[string]interface{}{
		"focused_activity": focusedActivity,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleWaitForActivity handles the wait_for_activity tool
func (h *ActivityHandler) handleWaitForActivity(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	activity, err := request.RequireString("activity")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	timeout := 10 // default 10 seconds
	if timeoutVal, ok := args["timeout"].(float64); ok {
		timeout = int(timeoutVal)
	}
	if timeout <= 0 {
		timeout = 10
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Activity.WaitForActivityToFocus(activity, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to wait for activity %s: %v", activity, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Activity %s is now focused", activity),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}