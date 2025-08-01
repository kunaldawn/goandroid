package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// InteractionHandler handles view interaction operations
type InteractionHandler struct {
	manager goandroid.AndroidManager
}

// NewInteractionHandler creates a new view interaction handler
func NewInteractionHandler(manager goandroid.AndroidManager) *InteractionHandler {
	return &InteractionHandler{manager: manager}
}

// RegisterTools registers view interaction tools
func (h *InteractionHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Click view tool
	clickViewTool := mcpLib.NewTool("click_view",
		mcpLib.WithDescription("Click on a UI element found by text, resource, or description"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("search_type",
			mcpLib.Required(),
			mcpLib.Description("Type of search: 'text', 'resource', 'description'"),
		),
		mcpLib.WithString("search_value",
			mcpLib.Required(),
			mcpLib.Description("Value to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element (0-based, default: 0)"),
		),
	)
	mcpServer.AddTool(clickViewTool, h.handleClickView)
}

// handleClickView handles the click_view tool
func (h *InteractionHandler) handleClickView(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	searchType, err := request.RequireString("search_type")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	searchValue, err := request.RequireString("search_value")
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

	var view interface{}
	var found bool

	switch searchType {
	case "text":
		view, found = views.GetByText(searchValue, index)
	case "resource":
		view, found = views.GetByResource(searchValue, index)
	case "description":
		view, found = views.GetByDescription(searchValue, index)
	default:
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid search_type: %s. Must be 'text', 'resource', or 'description'", searchType)), nil
	}

	if !found {
		return mcpLib.NewToolResultError(fmt.Sprintf("View with %s '%s' at index %d not found", searchType, searchValue, index)), nil
	}

	// Extract view from interface and get center coordinates
	viewStruct := view.(interface{})
	viewJSON, _ := json.Marshal(viewStruct)
	var viewMap map[string]interface{}
	json.Unmarshal(viewJSON, &viewMap)

	center := viewMap["Center"].(map[string]interface{})
	x := int(center["X"].(float64))
	y := int(center["Y"].(float64))

	err = device.Input.TouchScreen.Tap(x, y)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to click view: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Clicked view with %s '%s' at coordinates (%d, %d)", searchType, searchValue, x, y),
		"view":    viewStruct,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}