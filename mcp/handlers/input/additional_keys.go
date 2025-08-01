package input

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AdditionalKeysHandler handles additional key operations not covered in shortcuts
type AdditionalKeysHandler struct {
	manager goandroid.AndroidManager
}

// NewAdditionalKeysHandler creates a new additional keys handler
func NewAdditionalKeysHandler(manager goandroid.AndroidManager) *AdditionalKeysHandler {
	return &AdditionalKeysHandler{manager: manager}
}

// RegisterTools registers additional key operation tools
func (h *AdditionalKeysHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Press Call key
	pressCallTool := mcpLib.NewTool("press_call",
		mcpLib.WithDescription("Press the Call key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressCallTool, h.handlePressCall)

	// Press End Call key
	pressEndCallTool := mcpLib.NewTool("press_endcall",
		mcpLib.WithDescription("Press the End Call key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressEndCallTool, h.handlePressEndCall)

	// Press Camera key
	pressCameraTool := mcpLib.NewTool("press_camera",
		mcpLib.WithDescription("Press the Camera key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressCameraTool, h.handlePressCamera)
}

func (h *AdditionalKeysHandler) handlePressCall(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Count  float64 `json:"count"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	count := int(args.Count)
	if count <= 0 {
		count = 1
	}

	if err := device.Input.Key.PressCall(count); err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Call key pressed",
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdditionalKeysHandler) handlePressEndCall(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Count  float64 `json:"count"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	count := int(args.Count)
	if count <= 0 {
		count = 1
	}

	if err := device.Input.Key.PressEndCall(count); err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "End Call key pressed",
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdditionalKeysHandler) handlePressCamera(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Count  float64 `json:"count"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	count := int(args.Count)
	if count <= 0 {
		count = 1
	}

	if err := device.Input.Key.PressCamera(count); err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Camera key pressed",
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}