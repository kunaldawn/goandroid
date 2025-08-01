package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ViewRetrievalHandler handles getting View objects for elements
type ViewRetrievalHandler struct {
	manager goandroid.AndroidManager
}

// NewViewRetrievalHandler creates a new view retrieval handler
func NewViewRetrievalHandler(manager goandroid.AndroidManager) *ViewRetrievalHandler {
	return &ViewRetrievalHandler{
		manager: manager,
	}
}

// RegisterTools registers all view retrieval tools
func (h *ViewRetrievalHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get view for text
	getViewForTextTool := mcpLib.NewTool("get_view_for_text",
		mcpLib.WithDescription("Get a UI element View object by its text content"),
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
	mcpServer.AddTool(getViewForTextTool, h.handleGetViewForText)

	// Get view for matching text
	getViewForMatchingTextTool := mcpLib.NewTool("get_view_for_matching_text",
		mcpLib.WithDescription("Get a UI element View object by matching text content (case-insensitive substring match)"),
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
	mcpServer.AddTool(getViewForMatchingTextTool, h.handleGetViewForMatchingText)

	// Get view for resource
	getViewForResourceTool := mcpLib.NewTool("get_view_for_resource",
		mcpLib.WithDescription("Get a UI element View object by its resource ID"),
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
	mcpServer.AddTool(getViewForResourceTool, h.handleGetViewForResource)

	// Get view for matching resource
	getViewForMatchingResourceTool := mcpLib.NewTool("get_view_for_matching_resource",
		mcpLib.WithDescription("Get a UI element View object by matching resource ID (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("The resource ID pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(getViewForMatchingResourceTool, h.handleGetViewForMatchingResource)

	// Get view for description
	getViewForDescriptionTool := mcpLib.NewTool("get_view_for_description",
		mcpLib.WithDescription("Get a UI element View object by its content description"),
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
	mcpServer.AddTool(getViewForDescriptionTool, h.handleGetViewForDescription)

	// Get view for matching description
	getViewForMatchingDescriptionTool := mcpLib.NewTool("get_view_for_matching_description",
		mcpLib.WithDescription("Get a UI element View object by matching content description (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("The content description pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(getViewForMatchingDescriptionTool, h.handleGetViewForMatchingDescription)

	// Get view for type
	getViewForTypeTool := mcpLib.NewTool("get_view_for_type",
		mcpLib.WithDescription("Get a UI element View object by its class type"),
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
	mcpServer.AddTool(getViewForTypeTool, h.handleGetViewForType)

	// Get view for matching type
	getViewForMatchingTypeTool := mcpLib.NewTool("get_view_for_matching_type",
		mcpLib.WithDescription("Get a UI element View object by matching class type (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("The class type pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(getViewForMatchingTypeTool, h.handleGetViewForMatchingType)
}

func (h *ViewRetrievalHandler) handleGetViewForText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForMatchingText(args.Text, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForMatchingResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForMatchingDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ViewRetrievalHandler) handleGetViewForMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	view, err := device.View.GetViewForMatchingType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	jsonBytes, _ := json.Marshal(view)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}