package view

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	goandroidview "github.com/kunaldawn/goandroid/view"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// StateHandler handles view state checking operations
type StateHandler struct {
	manager goandroid.AndroidManager
}

// NewStateHandler creates a new state handler
func NewStateHandler(manager goandroid.AndroidManager) *StateHandler {
	return &StateHandler{
		manager: manager,
	}
}

// RegisterTools registers all state-checking tools
func (h *StateHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get view bounds
	getViewBoundsTool := mcpLib.NewTool("get_view_bounds",
		mcpLib.WithDescription("Get the bounds/rectangle of a view element"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("search_type",
			mcpLib.Required(),
			mcpLib.Description("Type of search: 'text', 'resource', 'description', or 'type'"),
		),
		mcpLib.WithString("search_value",
			mcpLib.Required(),
			mcpLib.Description("The value to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
	)
	mcpServer.AddTool(getViewBoundsTool, h.handleGetViewBounds)

	// Get view center
	getViewCenterTool := mcpLib.NewTool("get_view_center",
		mcpLib.WithDescription("Get the center point of a view element"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("search_type",
			mcpLib.Required(),
			mcpLib.Description("Type of search: 'text', 'resource', 'description', or 'type'"),
		),
		mcpLib.WithString("search_value",
			mcpLib.Required(),
			mcpLib.Description("The value to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
	)
	mcpServer.AddTool(getViewCenterTool, h.handleGetViewCenter)

	// Check if view is clickable
	isViewClickableTool := mcpLib.NewTool("is_view_clickable",
		mcpLib.WithDescription("Check if a view element is clickable"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("search_type",
			mcpLib.Required(),
			mcpLib.Description("Type of search: 'text', 'resource', 'description', or 'type'"),
		),
		mcpLib.WithString("search_value",
			mcpLib.Required(),
			mcpLib.Description("The value to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
	)
	mcpServer.AddTool(isViewClickableTool, h.handleIsViewClickable)

	// Get view info (comprehensive information)
	getViewInfoTool := mcpLib.NewTool("get_view_info",
		mcpLib.WithDescription("Get comprehensive information about a view element"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("search_type",
			mcpLib.Required(),
			mcpLib.Description("Type of search: 'text', 'resource', 'description', or 'type'"),
		),
		mcpLib.WithString("search_value",
			mcpLib.Required(),
			mcpLib.Description("The value to search for"),
		),
		mcpLib.WithNumber("index",
			mcpLib.Description("Index of the element if multiple matches (default: 0)"),
		),
	)
	mcpServer.AddTool(getViewInfoTool, h.handleGetViewInfo)

	// Get all views
	getAllViewsTool := mcpLib.NewTool("get_all_views",
		mcpLib.WithDescription("Get all views as a structured list with optional filtering"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithBoolean("clickable_only",
			mcpLib.Description("Only return clickable views (default: false)"),
		),
		mcpLib.WithBoolean("with_text_only",
			mcpLib.Description("Only return views with text (default: false)"),
		),
	)
	mcpServer.AddTool(getAllViewsTool, h.handleGetAllViews)
}

func (h *StateHandler) handleGetViewBounds(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		SearchType  string  `json:"search_type"`
		SearchValue string  `json:"search_value"`
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
	var view goandroidview.View
	var found bool

	switch args.SearchType {
	case "text":
		view, found = views.GetByText(args.SearchValue, index)
	case "resource":
		view, found = views.GetByResource(args.SearchValue, index)
	case "description":
		view, found = views.GetByDescription(args.SearchValue, index)
	case "type":
		view, found = views.GetByType(args.SearchValue, index)
	default:
		return mcpLib.NewToolResultError("Invalid search type"), nil
	}

	if !found {
		return mcpLib.NewToolResultError("View not found"), nil
	}

	data := map[string]interface{}{
		"bounds": map[string]interface{}{
			"top_left": map[string]interface{}{
				"x": view.Bound.TopLeft.X,
				"y": view.Bound.TopLeft.Y,
			},
			"bottom_right": map[string]interface{}{
				"x": view.Bound.BottomRight.X,
				"y": view.Bound.BottomRight.Y,
			},
		},
		"search_type":  args.SearchType,
		"search_value": args.SearchValue,
		"index":        index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *StateHandler) handleGetViewCenter(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		SearchType  string  `json:"search_type"`
		SearchValue string  `json:"search_value"`
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
	var view goandroidview.View
	var found bool

	switch args.SearchType {
	case "text":
		view, found = views.GetByText(args.SearchValue, index)
	case "resource":
		view, found = views.GetByResource(args.SearchValue, index)
	case "description":
		view, found = views.GetByDescription(args.SearchValue, index)
	case "type":
		view, found = views.GetByType(args.SearchValue, index)
	default:
		return mcpLib.NewToolResultError("Invalid search type"), nil
	}

	if !found {
		return mcpLib.NewToolResultError("View not found"), nil
	}

	data := map[string]interface{}{
		"center": map[string]interface{}{
			"x": view.Center.X,
			"y": view.Center.Y,
		},
		"search_type":  args.SearchType,
		"search_value": args.SearchValue,
		"index":        index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *StateHandler) handleIsViewClickable(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		SearchType  string  `json:"search_type"`
		SearchValue string  `json:"search_value"`
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
	var view goandroidview.View
	var found bool

	switch args.SearchType {
	case "text":
		view, found = views.GetByText(args.SearchValue, index)
	case "resource":
		view, found = views.GetByResource(args.SearchValue, index)
	case "description":
		view, found = views.GetByDescription(args.SearchValue, index)
	case "type":
		view, found = views.GetByType(args.SearchValue, index)
	default:
		return mcpLib.NewToolResultError("Invalid search type"), nil
	}

	if !found {
		return mcpLib.NewToolResultError("View not found"), nil
	}

	data := map[string]interface{}{
		"clickable":    view.Clickable,
		"search_type":  args.SearchType,
		"search_value": args.SearchValue,
		"index":        index,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *StateHandler) handleGetViewInfo(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial      string  `json:"serial"`
		SearchType  string  `json:"search_type"`
		SearchValue string  `json:"search_value"`
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
	var view goandroidview.View
	var found bool

	switch args.SearchType {
	case "text":
		view, found = views.GetByText(args.SearchValue, index)
	case "resource":
		view, found = views.GetByResource(args.SearchValue, index)
	case "description":
		view, found = views.GetByDescription(args.SearchValue, index)
	case "type":
		view, found = views.GetByType(args.SearchValue, index)
	default:
		return mcpLib.NewToolResultError("Invalid search type"), nil
	}

	if !found {
		return mcpLib.NewToolResultError("View not found"), nil
	}

	data := map[string]interface{}{
		"index":       view.Index,
		"class":       view.Class,
		"package":     view.Package,
		"resource":    view.Resource,
		"text":        view.Text,
		"description": view.Description,
		"clickable":   view.Clickable,
		"checked":     view.Checked,
		"enabled":     view.Enabled,
		"selected":    view.Selected,
		"bounds": map[string]interface{}{
			"top_left": map[string]interface{}{
				"x": view.Bound.TopLeft.X,
				"y": view.Bound.TopLeft.Y,
			},
			"bottom_right": map[string]interface{}{
				"x": view.Bound.BottomRight.X,
				"y": view.Bound.BottomRight.Y,
			},
		},
		"center": map[string]interface{}{
			"x": view.Center.X,
			"y": view.Center.Y,
		},
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *StateHandler) handleGetAllViews(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial        string `json:"serial"`
		ClickableOnly bool   `json:"clickable_only"`
		WithTextOnly  bool   `json:"with_text_only"`
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

	var viewList []map[string]interface{}
	for _, view := range views {
		// Apply filters
		if args.ClickableOnly && !view.Clickable {
			continue
		}
		if args.WithTextOnly && view.Text == "" {
			continue
		}

		viewData := map[string]interface{}{
			"index":       view.Index,
			"class":       view.Class,
			"resource":    view.Resource,
			"text":        view.Text,
			"description": view.Description,
			"clickable":   view.Clickable,
			"center": map[string]interface{}{
				"x": view.Center.X,
				"y": view.Center.Y,
			},
		}
		viewList = append(viewList, viewData)
	}

	data := map[string]interface{}{
		"count": len(viewList),
		"views": viewList,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}