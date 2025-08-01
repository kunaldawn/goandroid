package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ClickHandler handles direct click operations with timeout
type ClickHandler struct {
	manager goandroid.AndroidManager
}

// NewClickHandler creates a new click handler
func NewClickHandler(manager goandroid.AndroidManager) *ClickHandler {
	return &ClickHandler{
		manager: manager,
	}
}

// RegisterTools registers all click-related tools
func (h *ClickHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Click on text
	clickTextTool := mcpLib.NewTool("click_text",
		mcpLib.WithDescription("Click on a UI element by its text content"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("The text of the element to click"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickTextTool, h.handleClickText)

	// Click on matching text
	clickMatchingTextTool := mcpLib.NewTool("click_matching_text",
		mcpLib.WithDescription("Click on a UI element by matching text content (case-insensitive)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("The text pattern to match"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickMatchingTextTool, h.handleClickMatchingText)

	// Click on resource
	clickResourceTool := mcpLib.NewTool("click_resource",
		mcpLib.WithDescription("Click on a UI element by its resource ID"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("The resource ID of the element to click"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickResourceTool, h.handleClickResource)

	// Click on description
	clickDescriptionTool := mcpLib.NewTool("click_description",
		mcpLib.WithDescription("Click on a UI element by its content description"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("The content description of the element to click"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickDescriptionTool, h.handleClickDescription)

	// Click on type
	clickTypeTool := mcpLib.NewTool("click_type",
		mcpLib.WithDescription("Click on a UI element by its class type"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("The class type of the element to click"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickTypeTool, h.handleClickType)
}

func (h *ClickHandler) handleClickText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Clicked on text element",
		"text":    args.Text,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ClickHandler) handleClickMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickMatchingText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Clicked on matching text element",
		"text":    args.Text,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ClickHandler) handleClickResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success":  true,
		"message":  "Clicked on resource element",
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ClickHandler) handleClickDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success":     true,
		"message":     "Clicked on description element",
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ClickHandler) handleClickType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Clicked on type element",
		"type":    args.Type,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}