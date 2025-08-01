package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// MatchingOperationsHandler handles case-insensitive matching operations
type MatchingOperationsHandler struct {
	manager goandroid.AndroidManager
}

// NewMatchingOperationsHandler creates a new matching operations handler
func NewMatchingOperationsHandler(manager goandroid.AndroidManager) *MatchingOperationsHandler {
	return &MatchingOperationsHandler{
		manager: manager,
	}
}

// RegisterTools registers all matching operation tools
func (h *MatchingOperationsHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Matching presence checks
	isMatchingResourcePresentTool := mcpLib.NewTool("is_matching_resource_present",
		mcpLib.WithDescription("Check if resource ID is present on the screen (case-insensitive substring match)"),
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
	mcpServer.AddTool(isMatchingResourcePresentTool, h.handleIsMatchingResourcePresent)

	isMatchingDescriptionPresentTool := mcpLib.NewTool("is_matching_description_present",
		mcpLib.WithDescription("Check if content description is present on the screen (case-insensitive substring match)"),
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
	mcpServer.AddTool(isMatchingDescriptionPresentTool, h.handleIsMatchingDescriptionPresent)

	isMatchingTypePresentTool := mcpLib.NewTool("is_matching_type_present",
		mcpLib.WithDescription("Check if class type is present on the screen (case-insensitive substring match)"),
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
	mcpServer.AddTool(isMatchingTypePresentTool, h.handleIsMatchingTypePresent)

	// Matching click operations
	clickMatchingResourceTool := mcpLib.NewTool("click_matching_resource",
		mcpLib.WithDescription("Click on a UI element by matching resource ID (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("The resource ID pattern to match"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickMatchingResourceTool, h.handleClickMatchingResource)

	clickMatchingDescriptionTool := mcpLib.NewTool("click_matching_description",
		mcpLib.WithDescription("Click on a UI element by matching content description (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("The content description pattern to match"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickMatchingDescriptionTool, h.handleClickMatchingDescription)

	clickMatchingTypeTool := mcpLib.NewTool("click_matching_type",
		mcpLib.WithDescription("Click on a UI element by matching class type (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("The class type pattern to match"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 10)"),
		),
	)
	mcpServer.AddTool(clickMatchingTypeTool, h.handleClickMatchingType)

	// Find view by matching patterns
	findViewByMatchingTextTool := mcpLib.NewTool("find_view_by_matching_text",
		mcpLib.WithDescription("Find a UI element by matching text content (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("Text pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByMatchingTextTool, h.handleFindViewByMatchingText)

	findViewByMatchingResourceTool := mcpLib.NewTool("find_view_by_matching_resource",
		mcpLib.WithDescription("Find a UI element by matching resource ID (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("Resource ID pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByMatchingResourceTool, h.handleFindViewByMatchingResource)

	findViewByMatchingDescriptionTool := mcpLib.NewTool("find_view_by_matching_description",
		mcpLib.WithDescription("Find a UI element by matching content description (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("Content description pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByMatchingDescriptionTool, h.handleFindViewByMatchingDescription)

	findViewByMatchingTypeTool := mcpLib.NewTool("find_view_by_matching_type",
		mcpLib.WithDescription("Find a UI element by matching class type (case-insensitive substring match)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("Class type pattern to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByMatchingTypeTool, h.handleFindViewByMatchingType)
}

func (h *MatchingOperationsHandler) handleIsMatchingResourcePresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.IsMatchingResourcePresnt(args.Resource, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present":  present,
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleIsMatchingDescriptionPresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.IsMatchingDescriptionPresnt(args.Description, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present":     present,
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleIsMatchingTypePresent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.IsMatchingTypePresnt(args.Type, index, timeout)
	present := err == nil

	data := map[string]interface{}{
		"present": present,
		"type":    args.Type,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleClickMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickMatchingResource(args.Resource, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success":  true,
		"message":  "Clicked on matching resource element",
		"resource": args.Resource,
		"index":    index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleClickMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickMatchingDescription(args.Description, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success":     true,
		"message":     "Clicked on matching description element",
		"description": args.Description,
		"index":       index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleClickMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.View.ClickMatchingType(args.Type, index, timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Clicked on matching type element",
		"type":    args.Type,
		"index":   index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleFindViewByMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Text   string  `json:"text"`
		Index  float64 `json:"index"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	index := int(args.Index)
	view, found := views.GetByMatchingText(args.Text, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching text '%s' at index %d not found", args.Text, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	jsonBytes, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleFindViewByMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial   string  `json:"serial"`
		Resource string  `json:"resource"`
		Index    float64 `json:"index"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	index := int(args.Index)
	view, found := views.GetByMatchingResource(args.Resource, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching resource '%s' at index %d not found", args.Resource, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	jsonBytes, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleFindViewByMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		Description string  `json:"description"`
		Index       float64 `json:"index"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	index := int(args.Index)
	view, found := views.GetByMatchingDescription(args.Description, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching description '%s' at index %d not found", args.Description, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	jsonBytes, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *MatchingOperationsHandler) handleFindViewByMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Type   string  `json:"type"`
		Index  float64 `json:"index"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	index := int(args.Index)
	view, found := views.GetByMatchingType(args.Type, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching type '%s' at index %d not found", args.Type, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	jsonBytes, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}