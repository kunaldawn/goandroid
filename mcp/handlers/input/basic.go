package input

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kunaldawn/goandroid"
	"github.com/kunaldawn/goandroid/input"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// BasicInputHandler handles basic input operations
type BasicInputHandler struct {
	manager goandroid.AndroidManager
}

// NewBasicInputHandler creates a new basic input handler
func NewBasicInputHandler(manager goandroid.AndroidManager) *BasicInputHandler {
	return &BasicInputHandler{manager: manager}
}

// RegisterTools registers basic input operation tools
func (h *BasicInputHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Tap tool
	tapTool := mcpLib.NewTool("tap",
		mcpLib.WithDescription("Tap at specific coordinates on the device screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithNumber("x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate for the tap"),
		),
		mcpLib.WithNumber("y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate for the tap"),
		),
	)
	mcpServer.AddTool(tapTool, h.handleTap)

	// Key press tool
	keyPressTool := mcpLib.NewTool("key_press",
		mcpLib.WithDescription("Press a key on the device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("key",
			mcpLib.Required(),
			mcpLib.Description("Key to press (e.g., KEYCODE_HOME, KEYCODE_BACK)"),
		),
	)
	mcpServer.AddTool(keyPressTool, h.handleKeyPress)

	// Text input tool
	textInputTool := mcpLib.NewTool("text_input",
		mcpLib.WithDescription("Input text on the device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("Text to input"),
		),
	)
	mcpServer.AddTool(textInputTool, h.handleTextInput)
}

// getKeyCode maps key names to key codes
func getKeyCode(key string) (int, error) {
	keyMap := map[string]int{
		"KEYCODE_MENU":        input.MENU_KEY,
		"KEYCODE_HOME":        input.HOME_KEY,
		"KEYCODE_BACK":        input.BACK_KEY,
		"KEYCODE_CALL":        input.CALL_KEY,
		"KEYCODE_ENDCALL":     input.ENDCALL_KEY,
		"KEYCODE_DPAD_UP":     input.UP_KEY,
		"KEYCODE_DPAD_DOWN":   input.DOWN_KEY,
		"KEYCODE_DPAD_LEFT":   input.LEFT_KEY,
		"KEYCODE_DPAD_RIGHT":  input.RIGHT_KEY,
		"KEYCODE_VOLUME_UP":   input.VOLUME_UP_KEY,
		"KEYCODE_VOLUME_DOWN": input.VOLUME_DOWN_KEY,
		"KEYCODE_POWER":       input.POWER_KEY,
		"KEYCODE_CAMERA":      input.CAMERA_KEY,
		"KEYCODE_ENTER":       input.ENTER_KEY,
		"KEYCODE_DEL":         input.DEL_KEY,
	}

	// Support both with and without KEYCODE_ prefix
	normalizedKey := strings.ToUpper(key)
	if !strings.HasPrefix(normalizedKey, "KEYCODE_") {
		normalizedKey = "KEYCODE_" + normalizedKey
	}

	if code, exists := keyMap[normalizedKey]; exists {
		return code, nil
	}

	// Try to parse as integer
	if code, err := strconv.Atoi(key); err == nil {
		return code, nil
	}

	return 0, fmt.Errorf("unknown key: %s", key)
}

// handleTap handles the tap tool
func (h *BasicInputHandler) handleTap(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	xVal, ok := args["x"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("x must be a number"), nil
	}
	x := int(xVal)

	yVal, ok := args["y"].(float64)
	if !ok {
		return mcpLib.NewToolResultError("y must be a number"), nil
	}
	y := int(yVal)

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.TouchScreen.Tap(x, y)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to tap at (%v, %v): %v", x, y, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Tapped at coordinates (%v, %v)", x, y),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleKeyPress handles the key_press tool
func (h *BasicInputHandler) handleKeyPress(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	key, err := request.RequireString("key")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	keyCode, err := getKeyCode(key)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid key: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	err = device.Input.Key.Press(keyCode, 1)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to press key %s: %v", key, err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Pressed key: %s", key),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleTextInput handles the text_input tool
func (h *BasicInputHandler) handleTextInput(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	text, err := request.RequireString("text")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	device.Input.TextInput.EnterText(text)

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Input text: %s", text),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}