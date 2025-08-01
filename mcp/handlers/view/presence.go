package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// PresenceHandler handles view presence checking operations
type PresenceHandler struct {
	manager goandroid.AndroidManager
}

// NewPresenceHandler creates a new presence handler
func NewPresenceHandler(manager goandroid.AndroidManager) *PresenceHandler {
	return &PresenceHandler{
		manager: manager,
	}
}

// RegisterTools registers all presence-checking tools
func (h *PresenceHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Check if text is present
	isTextPresentTool := mcpLib.NewTool("is_text_present",
		mcpLib.WithDescription("Check if text is present on the screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("The text to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(isTextPresentTool, h.handleIsTextPresent)

	// Check if matching text is present
	isMatchingTextPresentTool := mcpLib.NewTool("is_matching_text_present",
		mcpLib.WithDescription("Check if matching text is present on the screen (case-insensitive)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("The text pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(isMatchingTextPresentTool, h.handleIsMatchingTextPresent)

	// Check if resource is present
	isResourcePresentTool := mcpLib.NewTool("is_resource_present",
		mcpLib.WithDescription("Check if resource ID is present on the screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("The resource ID to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(isResourcePresentTool, h.handleIsResourcePresent)

	// Check if description is present
	isDescriptionPresentTool := mcpLib.NewTool("is_description_present",
		mcpLib.WithDescription("Check if content description is present on the screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("The content description to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(isDescriptionPresentTool, h.handleIsDescriptionPresent)

	// Check if type is present
	isTypePresentTool := mcpLib.NewTool("is_type_present",
		mcpLib.WithDescription("Check if class type is present on the screen"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("The class type to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(isTypePresentTool, h.handleIsTypePresent)
}

func (h *PresenceHandler) handleIsTextPresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial  string  `json:"serial"`
		Text    string  `json:"text"`
		Index   float64 `json:"index"`
		Timeout float64 `json:"timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	index := int(args.Index)
	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 10
	}

	err := device.View.IsTextPresent(args.Text, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present": present,
		"text":    args.Text,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *PresenceHandler) handleIsMatchingTextPresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial  string  `json:"serial"`
		Text    string  `json:"text"`
		Index   float64 `json:"index"`
		Timeout float64 `json:"timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	index := int(args.Index)
	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 10
	}

	// Note: There's a typo in the original goandroid library method name
	err := device.View.IsMatchingTextPresnt(args.Text, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present": present,
		"text":    args.Text,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *PresenceHandler) handleIsResourcePresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial   string  `json:"serial"`
		Resource string  `json:"resource"`
		Index    float64 `json:"index"`
		Timeout  float64 `json:"timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	index := int(args.Index)
	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 10
	}

	err := device.View.IsResourcePresent(args.Resource, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present":  present,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *PresenceHandler) handleIsDescriptionPresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		Description string  `json:"description"`
		Index       float64 `json:"index"`
		Timeout     float64 `json:"timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	index := int(args.Index)
	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 10
	}

	err := device.View.IsDescriptionPresent(args.Description, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present":     present,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *PresenceHandler) handleIsTypePresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial  string  `json:"serial"`
		Type    string  `json:"type"`
		Index   float64 `json:"index"`
		Timeout float64 `json:"timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	index := int(args.Index)
	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 10
	}

	err := device.View.IsTypePresent(args.Type, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present": present,
		"type":    args.Type,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}