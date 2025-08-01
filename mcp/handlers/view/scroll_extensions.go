package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ScrollExtensionsHandler handles additional scroll operations not covered in main scrolling handler
type ScrollExtensionsHandler struct {
	manager goandroid.AndroidManager
}

// NewScrollExtensionsHandler creates a new scroll extensions handler
func NewScrollExtensionsHandler(manager goandroid.AndroidManager) *ScrollExtensionsHandler {
	return &ScrollExtensionsHandler{
		manager: manager,
	}
}

// RegisterTools registers all additional scroll tools
func (h *ScrollExtensionsHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Scroll up to resource
	scrollUpToResourceTool := mcpLib.NewTool("scroll_up_to_resource",
		mcpLib.WithDescription("Scroll up to find a UI element by resource ID"),
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
	mcpServer.AddTool(scrollUpToResourceTool, h.handleScrollUpToResource)

	// Scroll up to description
	scrollUpToDescriptionTool := mcpLib.NewTool("scroll_up_to_description",
		mcpLib.WithDescription("Scroll up to find a UI element by content description"),
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
	mcpServer.AddTool(scrollUpToDescriptionTool, h.handleScrollUpToDescription)

	// Scroll up to type
	scrollUpToTypeTool := mcpLib.NewTool("scroll_up_to_type",
		mcpLib.WithDescription("Scroll up to find a UI element by class type"),
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
	mcpServer.AddTool(scrollUpToTypeTool, h.handleScrollUpToType)

	// Scroll up to matching text
	scrollUpToMatchingTextTool := mcpLib.NewTool("scroll_up_to_matching_text",
		mcpLib.WithDescription("Scroll up to find a UI element by matching text (case-insensitive)"),
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
	mcpServer.AddTool(scrollUpToMatchingTextTool, h.handleScrollUpToMatchingText)

	// Scroll down to matching resource
	scrollDownToMatchingResourceTool := mcpLib.NewTool("scroll_down_to_matching_resource",
		mcpLib.WithDescription("Scroll down to find a UI element by matching resource ID (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToMatchingResourceTool, h.handleScrollDownToMatchingResource)

	// Scroll up to matching resource
	scrollUpToMatchingResourceTool := mcpLib.NewTool("scroll_up_to_matching_resource",
		mcpLib.WithDescription("Scroll up to find a UI element by matching resource ID (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollUpToMatchingResourceTool, h.handleScrollUpToMatchingResource)

	// Scroll down to matching description
	scrollDownToMatchingDescriptionTool := mcpLib.NewTool("scroll_down_to_matching_description",
		mcpLib.WithDescription("Scroll down to find a UI element by matching content description (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToMatchingDescriptionTool, h.handleScrollDownToMatchingDescription)

	// Scroll up to matching description
	scrollUpToMatchingDescriptionTool := mcpLib.NewTool("scroll_up_to_matching_description",
		mcpLib.WithDescription("Scroll up to find a UI element by matching content description (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollUpToMatchingDescriptionTool, h.handleScrollUpToMatchingDescription)

	// Scroll down to matching type
	scrollDownToMatchingTypeTool := mcpLib.NewTool("scroll_down_to_matching_type",
		mcpLib.WithDescription("Scroll down to find a UI element by matching class type (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollDownToMatchingTypeTool, h.handleScrollDownToMatchingType)

	// Scroll up to matching type
	scrollUpToMatchingTypeTool := mcpLib.NewTool("scroll_up_to_matching_type",
		mcpLib.WithDescription("Scroll up to find a UI element by matching class type (case-insensitive)"),
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
		mcpLib.WithNumber("max_scrolls",
			mcpLib.Description("Maximum number of scrolls (default: 10)"),
		),
	)
	mcpServer.AddTool(scrollUpToMatchingTypeTool, h.handleScrollUpToMatchingType)
}

func (h *ScrollExtensionsHandler) handleScrollUpToResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToResource(args.Resource, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":    found,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToDescription(args.Description, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":       found,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToType(args.Type, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found": found,
		"type":  args.Type,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToMatchingText(args.Text, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found": found,
		"text":  args.Text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollDownToMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollDownToMatchingResource(args.Resource, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":    found,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToMatchingResource(args.Resource, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":    found,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollDownToMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollDownToMatchingDescription(args.Description, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":       found,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToMatchingDescription(args.Description, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found":       found,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollDownToMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollDownToMatchingType(args.Type, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found": found,
		"type":  args.Type,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ScrollExtensionsHandler) handleScrollUpToMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ScrollUpToMatchingType(args.Type, index, maxScrolls)
	found := err == nil

	data := map[string]interface{}{
		"found": found,
		"type":  args.Type,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}