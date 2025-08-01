package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ScrollingHandler handles scroll operations to find elements
type ScrollingHandler struct {
	manager goandroid.AndroidManager
}

// NewScrollingHandler creates a new scrolling handler
func NewScrollingHandler(manager goandroid.AndroidManager) *ScrollingHandler {
	return &ScrollingHandler{
		manager: manager,
	}
}

// RegisterTools registers all scrolling-related tools
func (h *ScrollingHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Scroll down to text
	scrollDownToTextTool := mcpLib.NewTool("scroll_down_to_text",
		mcpLib.WithDescription("Scroll down to find a UI element by text"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToTextTool, h.handleScrollDownToText)

	// Scroll up to text
	scrollUpToTextTool := mcpLib.NewTool("scroll_up_to_text",
		mcpLib.WithDescription("Scroll up to find a UI element by text"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollUpToTextTool, h.handleScrollUpToText)

	// Scroll down to resource
	scrollDownToResourceTool := mcpLib.NewTool("scroll_down_to_resource",
		mcpLib.WithDescription("Scroll down to find a UI element by resource ID"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToResourceTool, h.handleScrollDownToResource)

	// Scroll down to description
	scrollDownToDescriptionTool := mcpLib.NewTool("scroll_down_to_description",
		mcpLib.WithDescription("Scroll down to find a UI element by content description"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToDescriptionTool, h.handleScrollDownToDescription)

	// Scroll down to type
	scrollDownToTypeTool := mcpLib.NewTool("scroll_down_to_type",
		mcpLib.WithDescription("Scroll down to find a UI element by class type"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToTypeTool, h.handleScrollDownToType)

	// Scroll down to matching text
	scrollDownToMatchingTextTool := mcpLib.NewTool("scroll_down_to_matching_text",
		mcpLib.WithDescription("Scroll down to find a UI element by matching text (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToMatchingTextTool, h.handleScrollDownToMatchingText)
}

func (h *ScrollingHandler) handleScrollDownToText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial     string  `json:"serial"`
		Text       string  `json:"text"`
		Index      float64 `json:"index"`
		MaxScrolls float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollDownToText(args.Text, index, maxScrolls)

	data := map[string]interface{}{
		"found": found,
		"text":  args.Text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollingHandler) handleScrollUpToText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial     string  `json:"serial"`
		Text       string  `json:"text"`
		Index      float64 `json:"index"`
		MaxScrolls float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollUpToText(args.Text, index, maxScrolls)

	data := map[string]interface{}{
		"found": found,
		"text":  args.Text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollingHandler) handleScrollDownToResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial     string  `json:"serial"`
		Resource   string  `json:"resource"`
		Index      float64 `json:"index"`
		MaxScrolls float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollDownToResource(args.Resource, index, maxScrolls)

	data := map[string]interface{}{
		"found":    found,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollingHandler) handleScrollDownToDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		Description string  `json:"description"`
		Index       float64 `json:"index"`
		MaxScrolls  float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollDownToDescription(args.Description, index, maxScrolls)

	data := map[string]interface{}{
		"found":       found,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollingHandler) handleScrollDownToType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial     string  `json:"serial"`
		Type       string  `json:"type"`
		Index      float64 `json:"index"`
		MaxScrolls float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollDownToType(args.Type, index, maxScrolls)

	data := map[string]interface{}{
		"found": found,
		"type":  args.Type,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollingHandler) handleScrollDownToMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial     string  `json:"serial"`
		Text       string  `json:"text"`
		Index      float64 `json:"index"`
		MaxScrolls float64 `json:"max_scrolls"`
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
	maxScrolls := int(args.MaxScrolls)
	if maxScrolls <= 0 {
		maxScrolls = 10
	}

	found := device.View.ScrollDownToMatchingText(args.Text, index, maxScrolls)

	data := map[string]interface{}{
		"found": found,
		"text":  args.Text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}