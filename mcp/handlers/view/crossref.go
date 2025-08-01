package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// CrossRefHandler handles cross-reference operations between view attributes
type CrossRefHandler struct {
	manager goandroid.AndroidManager
}

// NewCrossRefHandler creates a new cross-reference handler
func NewCrossRefHandler(manager goandroid.AndroidManager) *CrossRefHandler {
	return &CrossRefHandler{
		manager: manager,
	}
}

// RegisterTools registers all cross-reference tools
func (h *CrossRefHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get text for resource
	getTextForResourceTool := mcpLib.NewTool("get_text_for_resource",
		mcpLib.WithDescription("Get the text content of an element found by resource ID"),
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
	mcpServer.AddTool(getTextForResourceTool, h.handleGetTextForResource)

	// Get resource for text
	getResourceForTextTool := mcpLib.NewTool("get_resource_for_text",
		mcpLib.WithDescription("Get the resource ID of an element found by text"),
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
	mcpServer.AddTool(getResourceForTextTool, h.handleGetResourceForText)

	// Get description for text
	getDescriptionForTextTool := mcpLib.NewTool("get_description_for_text",
		mcpLib.WithDescription("Get the content description of an element found by text"),
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
	mcpServer.AddTool(getDescriptionForTextTool, h.handleGetDescriptionForText)

	// Get type for text
	getTypeForTextTool := mcpLib.NewTool("get_type_for_text",
		mcpLib.WithDescription("Get the class type of an element found by text"),
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
	mcpServer.AddTool(getTypeForTextTool, h.handleGetTypeForText)
}

func (h *CrossRefHandler) handleGetTextForResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"resource": args.Resource,
		"text":     text,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefHandler) handleGetResourceForText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	resource, err := device.View.GetResourceForText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"text":     args.Text,
		"resource": resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefHandler) handleGetDescriptionForText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	description, err := device.View.GetDescriptionForText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"text":        args.Text,
		"description": description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefHandler) handleGetTypeForText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	typename, err := device.View.GetTypeForText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"text":  args.Text,
		"type":  typename,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}