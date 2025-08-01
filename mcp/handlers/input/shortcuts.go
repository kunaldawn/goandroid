package input

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ShortcutsHandler handles common key shortcut operations
type ShortcutsHandler struct {
	manager goandroid.AndroidManager
}

// NewShortcutsHandler creates a new shortcuts handler
func NewShortcutsHandler(manager goandroid.AndroidManager) *ShortcutsHandler {
	return &ShortcutsHandler{
		manager: manager,
	}
}

// RegisterTools registers all shortcut-related tools
func (h *ShortcutsHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Press Home key
	pressHomeTool := mcpLib.NewTool("press_home",
		mcpLib.WithDescription("Press the Home key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressHomeTool, h.handlePressHome)

	// Press Back key
	pressBackTool := mcpLib.NewTool("press_back",
		mcpLib.WithDescription("Press the Back key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressBackTool, h.handlePressBack)

	// Press Menu key
	pressMenuTool := mcpLib.NewTool("press_menu",
		mcpLib.WithDescription("Press the Menu key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressMenuTool, h.handlePressMenu)

	// Press Power key
	pressPowerTool := mcpLib.NewTool("press_power",
		mcpLib.WithDescription("Press the Power key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressPowerTool, h.handlePressPower)

	// Press Enter key
	pressEnterTool := mcpLib.NewTool("press_enter",
		mcpLib.WithDescription("Press the Enter key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressEnterTool, h.handlePressEnter)

	// Press Delete key
	pressDeleteTool := mcpLib.NewTool("press_delete",
		mcpLib.WithDescription("Press the Delete/Backspace key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressDeleteTool, h.handlePressDelete)

	// Press Volume Up key
	pressVolumeUpTool := mcpLib.NewTool("press_volume_up",
		mcpLib.WithDescription("Press the Volume Up key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressVolumeUpTool, h.handlePressVolumeUp)

	// Press Volume Down key
	pressVolumeDownTool := mcpLib.NewTool("press_volume_down",
		mcpLib.WithDescription("Press the Volume Down key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressVolumeDownTool, h.handlePressVolumeDown)

	// Directional keys
	pressUpTool := mcpLib.NewTool("press_up",
		mcpLib.WithDescription("Press the Up directional key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressUpTool, h.handlePressUp)

	pressDownTool := mcpLib.NewTool("press_down",
		mcpLib.WithDescription("Press the Down directional key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressDownTool, h.handlePressDown)

	pressLeftTool := mcpLib.NewTool("press_left",
		mcpLib.WithDescription("Press the Left directional key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressLeftTool, h.handlePressLeft)

	pressRightTool := mcpLib.NewTool("press_right",
		mcpLib.WithDescription("Press the Right directional key on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("count",
			mcpLib.Description("Number of times to press the key (default: 1)"),
		),
	)
	mcpServer.AddTool(pressRightTool, h.handlePressRight)
}

func (h *ShortcutsHandler) handlePressHome(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressHome(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Home key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressBack(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressBack(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Back key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressMenu(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressMenu(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Menu key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressPower(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressPower(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Power key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressEnter(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressEnter(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Enter key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressDelete(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressDelete(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Delete key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressVolumeUp(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressVolumeUp(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Volume Up key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressVolumeDown(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressVolumeDown(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Volume Down key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressUp(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressUp(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Up key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressDown(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressDown(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Down key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressLeft(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressLeft(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Left key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ShortcutsHandler) handlePressRight(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Input.Key.PressRight(count)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Right key pressed",
		"serial":  args.Serial,
		"count":   count,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}