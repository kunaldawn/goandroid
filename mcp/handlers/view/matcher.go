package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// MatcherHandler handles matching view operations (case-insensitive)
type MatcherHandler struct {
	manager goandroid.AndroidManager
}

// NewMatcherHandler creates a new view matcher handler
func NewMatcherHandler(manager goandroid.AndroidManager) *MatcherHandler {
	return &MatcherHandler{manager: manager}
}

// RegisterTools registers matching view tools
func (h *MatcherHandler) RegisterTools(mcpServer *server.MCPServer) {
	tools := []struct {
		name        string
		description string
		paramName   string
		paramDesc   string
		handler     func(context.Context, mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error)
	}{
		{"find_view_by_matching_text", "Find a UI element by matching text content (case-insensitive)", "text", "Text pattern to search for", h.handleFindViewByMatchingText},
		{"find_view_by_matching_resource", "Find a UI element by matching resource ID (case-insensitive)", "resource", "Resource ID pattern to search for", h.handleFindViewByMatchingResource},
		{"find_view_by_matching_description", "Find a UI element by matching content description (case-insensitive)", "description", "Description pattern to search for", h.handleFindViewByMatchingDescription},
		{"find_view_by_matching_type", "Find a UI element by matching class type (case-insensitive)", "type", "Class type pattern to search for", h.handleFindViewByMatchingType},
	}

	for _, tool := range tools {
		mcpTool := mcpLib.NewTool(tool.name,
			mcpLib.WithDescription(tool.description),
			mcpLib.WithString("serial",
				mcpLib.Required(),
				mcpLib.Description("Serial number of the Android device"),
			),
			mcpLib.WithString(tool.paramName,
				mcpLib.Required(),
				mcpLib.Description(tool.paramDesc),
			),
			mcpLib.WithNumber("index",
				mcpLib.Description("Index of the element (0-based, default: 0)"),
			),
		)
		mcpServer.AddTool(mcpTool, tool.handler)
	}
}

func (h *MatcherHandler) getViewArgs(request mcpLib.CallToolRequest) (string, string, int, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return "", "", 0, err
	}

	// Get the search value from any of the possible parameter names
	args := request.GetArguments()
	var searchValue string
	for _, key := range []string{"text", "resource", "description", "type"} {
		if val, ok := args[key].(string); ok {
			searchValue = val
			break
		}
	}
	if searchValue == "" {
		return "", "", 0, fmt.Errorf("search parameter not found")
	}

	index := 0
	if indexVal, ok := args["index"].(float64); ok {
		index = int(indexVal)
	}
	if index < 0 {
		index = 0
	}

	return serial, searchValue, index, nil
}

// handleFindViewByMatchingText handles the find_view_by_matching_text tool
func (h *MatcherHandler) handleFindViewByMatchingText(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, text, index, err := h.getViewArgs(request)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByMatchingText(text, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching text '%s' at index %d not found", text, index)), nil
	}

	result := map[string]interface{}{"found": true, "view": view}
	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByMatchingResource handles the find_view_by_matching_resource tool
func (h *MatcherHandler) handleFindViewByMatchingResource(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, resource, index, err := h.getViewArgs(request)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByMatchingResource(resource, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching resource '%s' at index %d not found", resource, index)), nil
	}

	result := map[string]interface{}{"found": true, "view": view}
	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByMatchingDescription handles the find_view_by_matching_description tool
func (h *MatcherHandler) handleFindViewByMatchingDescription(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, description, index, err := h.getViewArgs(request)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByMatchingDescription(description, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching description '%s' at index %d not found", description, index)), nil
	}

	result := map[string]interface{}{"found": true, "view": view}
	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleFindViewByMatchingType handles the find_view_by_matching_type tool
func (h *MatcherHandler) handleFindViewByMatchingType(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, viewType, index, err := h.getViewArgs(request)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	views, err := device.View.GetViewes()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get views: %v", err)), nil
	}

	view, found := views.GetByMatchingType(viewType, index)
	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with matching type '%s' at index %d not found", viewType, index)), nil
	}

	result := map[string]interface{}{"found": true, "view": view}
	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}