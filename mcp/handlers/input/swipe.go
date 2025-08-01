package input

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// SwipeHandler handles swipe input operations
type SwipeHandler struct {
	manager goandroid.AndroidManager
}

// NewSwipeHandler creates a new swipe handler
func NewSwipeHandler(manager goandroid.AndroidManager) *SwipeHandler {
	return &SwipeHandler{manager: manager}
}

// RegisterTools registers swipe operation tools
func (h *SwipeHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Swipe tool
	swipeTool := mcpLib.NewTool("swipe",
		mcpLib.WithDescription("Swipe from one coordinate to another"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("x1",
			mcpLib.Required(),
			mcpLib.Description("Starting X coordinate"),
		),
		mcpLib.WithNumber("y1",
			mcpLib.Required(),
			mcpLib.Description("Starting Y coordinate"),
		),
		mcpLib.WithNumber("x2",
			mcpLib.Required(),
			mcpLib.Description("Ending X coordinate"),
		),
		mcpLib.WithNumber("y2",
			mcpLib.Required(),
			mcpLib.Description("Ending Y coordinate"),
		),
		mcpLib.WithNumber("delay",
			mcpLib.Description("Delay in milliseconds (default: 300)"),
		),
	)
	mcpServer.AddTool(swipeTool, h.handleSwipe)

	// Swipe up tool
	swipeUpTool := mcpLib.NewTool("swipe_up",
		mcpLib.WithDescription("Swipe up on the device screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of swipes (default: 1)"),
		),
	)
	mcpServer.AddTool(swipeUpTool, h.handleSwipeUp)

	// Swipe down tool
	swipeDownTool := mcpLib.NewTool("swipe_down",
		mcpLib.WithDescription("Swipe down on the device screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of swipes (default: 1)"),
		),
	)
	mcpServer.AddTool(swipeDownTool, h.handleSwipeDown)

	// Swipe left tool
	swipeLeftTool := mcpLib.NewTool("swipe_left",
		mcpLib.WithDescription("Swipe left on the device screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of swipes (default: 1)"),
		),
	)
	mcpServer.AddTool(swipeLeftTool, h.handleSwipeLeft)

	// Swipe right tool
	swipeRightTool := mcpLib.NewTool("swipe_right",
		mcpLib.WithDescription("Swipe right on the device screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of swipes (default: 1)"),
		),
	)
	mcpServer.AddTool(swipeRightTool, h.handleSwipeRight)
}

// handleSwipe handles the swipe tool
func (h *SwipeHandler) handleSwipe(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	x1Val, ok := args["x1"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("x1 must be a number"), nil
	}
	x1 := int(x1Val)

	y1Val, ok := args["y1"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("y1 must be a number"), nil
	}
	y1 := int(y1Val)

	x2Val, ok := args["x2"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("x2 must be a number"), nil
	}
	x2 := int(x2Val)

	y2Val, ok := args["y2"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("y2 must be a number"), nil
	}
	y2 := int(y2Val)

	delay := 300 // default delay
	if delayVal, ok := args["delay"].(float64); ok {
		delay = int(delayVal)
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.Swipe(x1, y1, x2, y2, delay)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to swipe from (%d, %d) to (%d, %d): %v", x1, y1, x2, y2, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Swiped from (%d, %d) to (%d, %d) with delay %dms", x1, y1, x2, y2, delay),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleSwipeUp handles the swipe_up tool
func (h *SwipeHandler) handleSwipeUp(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	count := 1
	if countVal, ok := args["count"].(float64); ok {
		count = int(countVal)
	}
	if count <= 0 {
		count = 1
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.SwipeUp(count)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to swipe up %d times: %v", count, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Swiped up %d times", count),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleSwipeDown handles the swipe_down tool
func (h *SwipeHandler) handleSwipeDown(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	count := 1
	if countVal, ok := args["count"].(float64); ok {
		count = int(countVal)
	}
	if count <= 0 {
		count = 1
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.SwipeDown(count)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to swipe down %d times: %v", count, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Swiped down %d times", count),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleSwipeLeft handles the swipe_left tool
func (h *SwipeHandler) handleSwipeLeft(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	count := 1
	if countVal, ok := args["count"].(float64); ok {
		count = int(countVal)
	}
	if count <= 0 {
		count = 1
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.SwipeLeft(count)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to swipe left %d times: %v", count, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Swiped left %d times", count),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleSwipeRight handles the swipe_right tool
func (h *SwipeHandler) handleSwipeRight(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	count := 1
	if countVal, ok := args["count"].(float64); ok {
		count = int(countVal)
	}
	if count <= 0 {
		count = 1
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.SwipeRight(count)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to swipe right %d times: %v", count, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Swiped right %d times", count),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}