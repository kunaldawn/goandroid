package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// CrossRefExtensionsHandler handles additional cross-reference operations
type CrossRefExtensionsHandler struct {
	manager goandroid.AndroidManager
}

// NewCrossRefExtensionsHandler creates a new cross-reference extensions handler
func NewCrossRefExtensionsHandler(manager goandroid.AndroidManager) *CrossRefExtensionsHandler {
	return &CrossRefExtensionsHandler{
		manager: manager,
	}
}

// RegisterTools registers all additional cross-reference tools
func (h *CrossRefExtensionsHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get text for type
	getTextForTypeTool := mcpLib.NewTool("get_text_for_type",
		mcpLib.WithDescription("Get the text content of an element found by class type"),
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
	mcpServer.AddTool(getTextForTypeTool, h.handleGetTextForType)

	// Get text for description
	getTextForDescriptionTool := mcpLib.NewTool("get_text_for_description",
		mcpLib.WithDescription("Get the text content of an element found by content description"),
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
	mcpServer.AddTool(getTextForDescriptionTool, h.handleGetTextForDescription)

	// Get resource for type
	getResourceForTypeTool := mcpLib.NewTool("get_resource_for_type",
		mcpLib.WithDescription("Get the resource ID of an element found by class type"),
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
	mcpServer.AddTool(getResourceForTypeTool, h.handleGetResourceForType)

	// Get resource for description
	getResourceForDescriptionTool := mcpLib.NewTool("get_resource_for_description",
		mcpLib.WithDescription("Get the resource ID of an element found by content description"),
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
	mcpServer.AddTool(getResourceForDescriptionTool, h.handleGetResourceForDescription)

	// Get description for resource
	getDescriptionForResourceTool := mcpLib.NewTool("get_description_for_resource",
		mcpLib.WithDescription("Get the content description of an element found by resource ID"),
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
	mcpServer.AddTool(getDescriptionForResourceTool, h.handleGetDescriptionForResource)

	// Get description for type
	getDescriptionForTypeTool := mcpLib.NewTool("get_description_for_type",
		mcpLib.WithDescription("Get the content description of an element found by class type"),
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
	mcpServer.AddTool(getDescriptionForTypeTool, h.handleGetDescriptionForType)

	// Get type for resource
	getTypeForResourceTool := mcpLib.NewTool("get_type_for_resource",
		mcpLib.WithDescription("Get the class type of an element found by resource ID"),
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
	mcpServer.AddTool(getTypeForResourceTool, h.handleGetTypeForResource)

	// Get type for description
	getTypeForDescriptionTool := mcpLib.NewTool("get_type_for_description",
		mcpLib.WithDescription("Get the class type of an element found by content description"),
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
	mcpServer.AddTool(getTypeForDescriptionTool, h.handleGetTypeForDescription)

	// Matching variants - Get text for matching type
	getTextForMatchingTypeTool := mcpLib.NewTool("get_text_for_matching_type",
		mcpLib.WithDescription("Get the text content of an element found by matching class type (case-insensitive)"),
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
	mcpServer.AddTool(getTextForMatchingTypeTool, h.handleGetTextForMatchingType)

	// Get text for matching description
	getTextForMatchingDescriptionTool := mcpLib.NewTool("get_text_for_matching_description",
		mcpLib.WithDescription("Get the text content of an element found by matching content description (case-insensitive)"),
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
	mcpServer.AddTool(getTextForMatchingDescriptionTool, h.handleGetTextForMatchingDescription)

	// Get text for matching resource
	getTextForMatchingResourceTool := mcpLib.NewTool("get_text_for_matching_resource",
		mcpLib.WithDescription("Get the text content of an element found by matching resource ID (case-insensitive)"),
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
	mcpServer.AddTool(getTextForMatchingResourceTool, h.handleGetTextForMatchingResource)

	// Get resource for matching type
	getResourceForMatchingTypeTool := mcpLib.NewTool("get_resource_for_matching_type",
		mcpLib.WithDescription("Get the resource ID of an element found by matching class type (case-insensitive)"),
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
	mcpServer.AddTool(getResourceForMatchingTypeTool, h.handleGetResourceForMatchingType)

	// Get resource for matching description
	getResourceForMatchingDescriptionTool := mcpLib.NewTool("get_resource_for_matching_description",
		mcpLib.WithDescription("Get the resource ID of an element found by matching content description (case-insensitive)"),
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
	mcpServer.AddTool(getResourceForMatchingDescriptionTool, h.handleGetResourceForMatchingDescription)
}

func (h *CrossRefExtensionsHandler) handleGetTextForType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"type":  args.Type,
		"text":  text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetTextForDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"description": args.Description,
		"text":        text,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetResourceForType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	resource, err := device.View.GetResourceForType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"type":     args.Type,
		"resource": resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetResourceForDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	resource, err := device.View.GetResourceForDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"description": args.Description,
		"resource":    resource,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetDescriptionForResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	description, err := device.View.GetDescriptionForResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"resource":    args.Resource,
		"description": description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetDescriptionForType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	description, err := device.View.GetDescriptionFortype(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"type":        args.Type,
		"description": description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetTypeForResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	typename, err := device.View.GetTypeForResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"resource": args.Resource,
		"type":     typename,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetTypeForDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	typename, err := device.View.GetTypeForDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"description": args.Description,
		"type":        typename,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

// Matching variants
func (h *CrossRefExtensionsHandler) handleGetTextForMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForMatchingType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"type":  args.Type,
		"text":  text,
		"index": index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetTextForMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForMatchingDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"description": args.Description,
		"text":        text,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetTextForMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	text, err := device.View.GetTextForMatchingResource(args.Resource, index, timeout)
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

func (h *CrossRefExtensionsHandler) handleGetResourceForMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	resource, err := device.View.GetResourceForMatchingType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"type":     args.Type,
		"resource": resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *CrossRefExtensionsHandler) handleGetResourceForMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	resource, err := device.View.GetResourceForMatchingDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"description": args.Description,
		"resource":    resource,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}