package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// FinderHandler handles view finding operations
type FinderHandler struct {
	manager goandroid.AndroidManager
}

// NewFinderHandler creates a new view finder handler
func NewFinderHandler(manager goandroid.AndroidManager) *FinderHandler {
	return &FinderHandler{manager: manager}
}

// RegisterTools registers view finding tools
func (h *FinderHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get view hierarchy tool
	getViewHierarchyTool := mcpLib.NewTool("get_view_hierarchy",
		mcpLib.WithDescription("Get the UI view hierarchy of the device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getViewHierarchyTool, h.handleGetViewHierarchy)

	// Find view by text tool
	findViewByTextTool := mcpLib.NewTool("find_view_by_text",
		mcpLib.WithDescription("Find a UI element by its text content"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("text",
			mcpLib.Required(),
			mcpLib.Description("Text to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByTextTool, h.handleFindViewByText)

	// Find view by resource tool
	findViewByResourceTool := mcpLib.NewTool("find_view_by_resource",
		mcpLib.WithDescription("Find a UI element by its resource ID"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("resource",
			mcpLib.Required(),
			mcpLib.Description("Resource ID to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByResourceTool, h.handleFindViewByResource)

	// Find view by description tool
	findViewByDescriptionTool := mcpLib.NewTool("find_view_by_description",
		mcpLib.WithDescription("Find a UI element by its content description"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("description",
			mcpLib.Required(),
			mcpLib.Description("Content description to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByDescriptionTool, h.handleFindViewByDescription)

	// Find view by type tool
	findViewByTypeTool := mcpLib.NewTool("find_view_by_type",
		mcpLib.WithDescription("Find a UI element by its class type"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("type",
			mcpLib.Required(),
			mcpLib.Description("Class type to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(findViewByTypeTool, h.handleFindViewByType)
}

// handleGetViewHierarchy handles the get_view_hierarchy tool
func (h *FinderHandler) handleGetViewHierarchy(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get view hierarchy: %v", err)), nil
	}

	result := map[string]interface{}{
		"views": views,
		"count": len(views),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByText handles the find_view_by_text tool
func (h *FinderHandler) handleFindViewByText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	text, err := request.RequireString("text")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	index := 0
	if indexVal, ok := args["index"].(float64); ok {
		index = int(indexVal)
	}
	if index < 0 {
		index = 0
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByText(text, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with text '%s' at index %d not found", text, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByResource handles the find_view_by_resource tool
func (h *FinderHandler) handleFindViewByResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	resource, err := request.RequireString("resource")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	index := 0
	if indexVal, ok := args["index"].(float64); ok {
		index = int(indexVal)
	}
	if index < 0 {
		index = 0
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByResource(resource, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with resource '%s' at index %d not found", resource, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByDescription handles the find_view_by_description tool
func (h *FinderHandler) handleFindViewByDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	description, err := request.RequireString("description")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	index := 0
	if indexVal, ok := args["index"].(float64); ok {
		index = int(indexVal)
	}
	if index < 0 {
		index = 0
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByDescription(description, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with description '%s' at index %d not found", description, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByType handles the find_view_by_type tool
func (h *FinderHandler) handleFindViewByType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	viewType, err := request.RequireString("type")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	index := 0
	if indexVal, ok := args["index"].(float64); ok {
		index = int(indexVal)
	}
	if index < 0 {
		index = 0
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByType(viewType, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with type '%s' at index %d not found", viewType, index)), nil
	}

	result := map[string]interface{}{
		"found": true,
		"view":  view,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}