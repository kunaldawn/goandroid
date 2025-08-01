package geometry

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// GeometryUtilsHandler handles geometric calculations and utilities
type GeometryUtilsHandler struct{}

// NewGeometryUtilsHandler creates a new geometry utils handler
func NewGeometryUtilsHandler() *GeometryUtilsHandler {
	return &GeometryUtilsHandler{}
}

// RegisterTools registers all geometry utility tools
func (h *GeometryUtilsHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Calculate distance between two points
	calculateDistanceTool := mcpLib.NewTool("calculate_distance",
		mcpLib.WithDescription("Calculate the distance between two points"),
		mcpLib.WithNumber("x1",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the first point"),
		),
		mcpLib.WithNumber("y1",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the first point"),
		),
		mcpLib.WithNumber("x2",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the second point"),
		),
		mcpLib.WithNumber("y2",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the second point"),
		),
	)
	mcpServer.AddTool(calculateDistanceTool, h.handleCalculateDistance)

	// Calculate rectangle center
	calculateRectCenterTool := mcpLib.NewTool("calculate_rect_center",
		mcpLib.WithDescription("Calculate the center point of a rectangle"),
		mcpLib.WithNumber("top_left_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the top-left corner"),
		),
		mcpLib.WithNumber("top_left_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the top-left corner"),
		),
		mcpLib.WithNumber("bottom_right_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the bottom-right corner"),
		),
		mcpLib.WithNumber("bottom_right_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the bottom-right corner"),
		),
	)
	mcpServer.AddTool(calculateRectCenterTool, h.handleCalculateRectCenter)

	// Calculate rectangle area
	calculateRectAreaTool := mcpLib.NewTool("calculate_rect_area",
		mcpLib.WithDescription("Calculate the area of a rectangle"),
		mcpLib.WithNumber("top_left_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the top-left corner"),
		),
		mcpLib.WithNumber("top_left_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the top-left corner"),
		),
		mcpLib.WithNumber("bottom_right_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the bottom-right corner"),
		),
		mcpLib.WithNumber("bottom_right_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the bottom-right corner"),
		),
	)
	mcpServer.AddTool(calculateRectAreaTool, h.handleCalculateRectArea)

	// Check if point is inside rectangle
	pointInRectTool := mcpLib.NewTool("point_in_rectangle",
		mcpLib.WithDescription("Check if a point is inside a rectangle"),
		mcpLib.WithNumber("point_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the point"),
		),
		mcpLib.WithNumber("point_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the point"),
		),
		mcpLib.WithNumber("rect_top_left_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the rectangle's top-left corner"),
		),
		mcpLib.WithNumber("rect_top_left_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the rectangle's top-left corner"),
		),
		mcpLib.WithNumber("rect_bottom_right_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the rectangle's bottom-right corner"),
		),
		mcpLib.WithNumber("rect_bottom_right_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the rectangle's bottom-right corner"),
		),
	)
	mcpServer.AddTool(pointInRectTool, h.handlePointInRectangle)

	// Create offset point
	createOffsetPointTool := mcpLib.NewTool("create_offset_point",
		mcpLib.WithDescription("Create a new point by adding an offset to an existing point"),
		mcpLib.WithNumber("original_x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate of the original point"),
		),
		mcpLib.WithNumber("original_y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate of the original point"),
		),
		mcpLib.WithNumber("offset_x",
			mcpLib.Required(),
			mcpLib.Description("X offset to add"),
		),
		mcpLib.WithNumber("offset_y",
			mcpLib.Required(),
			mcpLib.Description("Y offset to add"),
		),
	)
	mcpServer.AddTool(createOffsetPointTool, h.handleCreateOffsetPoint)
}

func (h *GeometryUtilsHandler) handleCalculateDistance(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		X1 float64 `json:"x1"`
		Y1 float64 `json:"y1"`
		X2 float64 `json:"x2"`
		Y2 float64 `json:"y2"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	// Calculate Euclidean distance
	dx := args.X2 - args.X1
	dy := args.Y2 - args.Y1
	distance := math.Sqrt(dx*dx + dy*dy)

	data := map[string]interface{}{
		"point1":   map[string]interface{}{"x": args.X1, "y": args.Y1},
		"point2":   map[string]interface{}{"x": args.X2, "y": args.Y2},
		"distance": distance,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *GeometryUtilsHandler) handleCalculateRectCenter(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		TopLeftX     float64 `json:"top_left_x"`
		TopLeftY     float64 `json:"top_left_y"`
		BottomRightX float64 `json:"bottom_right_x"`
		BottomRightY float64 `json:"bottom_right_y"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	centerX := (args.TopLeftX + args.BottomRightX) / 2
	centerY := (args.TopLeftY + args.BottomRightY) / 2

	data := map[string]interface{}{
		"rectangle": map[string]interface{}{
			"top_left":     map[string]interface{}{"x": args.TopLeftX, "y": args.TopLeftY},
			"bottom_right": map[string]interface{}{"x": args.BottomRightX, "y": args.BottomRightY},
		},
		"center": map[string]interface{}{"x": centerX, "y": centerY},
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *GeometryUtilsHandler) handleCalculateRectArea(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		TopLeftX     float64 `json:"top_left_x"`
		TopLeftY     float64 `json:"top_left_y"`
		BottomRightX float64 `json:"bottom_right_x"`
		BottomRightY float64 `json:"bottom_right_y"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	width := math.Abs(args.BottomRightX - args.TopLeftX)
	height := math.Abs(args.BottomRightY - args.TopLeftY)
	area := width * height

	data := map[string]interface{}{
		"rectangle": map[string]interface{}{
			"top_left":     map[string]interface{}{"x": args.TopLeftX, "y": args.TopLeftY},
			"bottom_right": map[string]interface{}{"x": args.BottomRightX, "y": args.BottomRightY},
		},
		"width":  width,
		"height": height,
		"area":   area,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *GeometryUtilsHandler) handlePointInRectangle(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		PointX         float64 `json:"point_x"`
		PointY         float64 `json:"point_y"`
		RectTopLeftX   float64 `json:"rect_top_left_x"`
		RectTopLeftY   float64 `json:"rect_top_left_y"`
		RectBottomRightX float64 `json:"rect_bottom_right_x"`
		RectBottomRightY float64 `json:"rect_bottom_right_y"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	// Check if point is inside rectangle
	inside := args.PointX >= args.RectTopLeftX && 
		      args.PointX <= args.RectBottomRightX &&
		      args.PointY >= args.RectTopLeftY && 
		      args.PointY <= args.RectBottomRightY

	data := map[string]interface{}{
		"point": map[string]interface{}{"x": args.PointX, "y": args.PointY},
		"rectangle": map[string]interface{}{
			"top_left":     map[string]interface{}{"x": args.RectTopLeftX, "y": args.RectTopLeftY},
			"bottom_right": map[string]interface{}{"x": args.RectBottomRightX, "y": args.RectBottomRightY},
		},
		"inside": inside,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *GeometryUtilsHandler) handleCreateOffsetPoint(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		OriginalX float64 `json:"original_x"`
		OriginalY float64 `json:"original_y"`
		OffsetX   float64 `json:"offset_x"`
		OffsetY   float64 `json:"offset_y"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	newX := args.OriginalX + args.OffsetX
	newY := args.OriginalY + args.OffsetY

	data := map[string]interface{}{
		"original_point": map[string]interface{}{"x": args.OriginalX, "y": args.OriginalY},
		"offset":         map[string]interface{}{"x": args.OffsetX, "y": args.OffsetY},
		"new_point":      map[string]interface{}{"x": newX, "y": newY},
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}