package input

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	"github.com/kunaldawn/goandroid/geometry"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// GestureHandler handles gesture input operations
type GestureHandler struct {
	manager goandroid.AndroidManager
}

// NewGestureHandler creates a new gesture handler
func NewGestureHandler(manager goandroid.AndroidManager) *GestureHandler {
	return &GestureHandler{manager: manager}
}

// RegisterTools registers gesture operation tools
func (h *GestureHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Draw gesture tool
	drawGestureTool := mcpLib.NewTool("draw_gesture",
		mcpLib.WithDescription("Draw a continuous gesture with multiple points"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithArray("points",
			mcpLib.Required(),
			mcpLib.Description("Array of points for the gesture, each point should have x and y properties"),
		),
		mcpLib.WithNumber("delay",
			mcpLib.Description("Delay between points in milliseconds (default: 50)"),
		),
		mcpLib.WithBoolean("emulator",
			mcpLib.Description("Use emulator gesture mode (default: false)"),
		),
	)
	mcpServer.AddTool(drawGestureTool, h.handleDrawGesture)

	// Draw gesture for emulator
	drawGestureEmulatorTool := mcpLib.NewTool("draw_gesture_emulator",
		mcpLib.WithDescription("Draw a continuous gesture on an emulator device using single-touch protocol"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithArray("points",
			mcpLib.Required(),
			mcpLib.Description("Array of points defining the gesture path. Each point should have 'x' and 'y' coordinates"),
		),
		mcpLib.WithNumber("delay",
			mcpLib.Description("Delay between points in milliseconds (default: 100)"),
		),
	)
	mcpServer.AddTool(drawGestureEmulatorTool, h.handleDrawGestureEmulator)
}

// handleDrawGesture handles the draw_gesture tool
func (h *GestureHandler) handleDrawGesture(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	pointsData, ok := args["points"].([]interface{})
	if !ok {
		return mcpLib.NewToolResultError("points must be an array"), nil
	}

	var points geometry.Points
	for _, pointData := range pointsData {
		pointMap, ok := pointData.(map[string]interface{})
		if !ok {
			return mcpLib.NewToolResultError("each point must be an object with x and y properties"), nil
		}

		xVal, ok := pointMap["x"].(float64)
		if !ok {
			return mcpLib.NewToolResultError("point x must be a number"), nil
		}

		yVal, ok := pointMap["y"].(float64)
		if !ok {
			return mcpLib.NewToolResultError("point y must be a number"), nil
		}

		points = append(points, geometry.Point{X: int(xVal), Y: int(yVal)})
	}

	if len(points) == 0 {
		return mcpLib.NewToolResultError("at least one point is required"), nil
	}

	delay := 50 // default delay
	if delayVal, ok := args["delay"].(float64); ok {
		delay = int(delayVal)
	}

	emulator := false
	if emulatorVal, ok := args["emulator"].(bool); ok {
		emulator = emulatorVal
	}

	device := h.manager.GetNewAndroidDevice(serial)
	var gestureErr error
	if emulator {
		gestureErr = device.Input.TouchScreen.DrawGestureEmulator(points, delay)
	} else {
		gestureErr = device.Input.TouchScreen.DrawGesture(points, delay)
	}

	if gestureErr != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to draw gesture: %v", gestureErr)), nil
	}

	result := map[string]interface{}{
		"success":  true,
		"message":  fmt.Sprintf("Drew gesture with %d points", len(points)),
		"points":   points,
		"delay":    delay,
		"emulator": emulator,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleDrawGestureEmulator handles the draw_gesture_emulator tool
func (h *GestureHandler) handleDrawGestureEmulator(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Points []map[string]interface{} `json:"points"`
		Delay  float64 `json:"delay"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	// Convert points to geometry.Points
	points := make(geometry.Points, len(args.Points))
	for i, point := range args.Points {
		x, xOk := point["x"].(float64)
		y, yOk := point["y"].(float64)
		if !xOk || !yOk {
			return mcpLib.NewToolResultError(fmt.Sprintf("Invalid point at index %d: points must have 'x' and 'y' coordinates", i)), nil
		}
		points[i] = geometry.Point{X: int(x), Y: int(y)}
	}

	if len(points) == 0 {
		return mcpLib.NewToolResultError("At least one point is required"), nil
	}

	delay := int(args.Delay)
	if delay <= 0 {
		delay = 100 // Default delay
	}

	// Always use emulator gesture method
	gestureErr := device.Input.TouchScreen.DrawGestureEmulator(points, delay)
	if gestureErr != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to draw emulator gesture: %v", gestureErr)), nil
	}

	result := map[string]interface{}{
		"success":     true,
		"points":      len(points),
		"delay":       delay,
		"device_type": "emulator",
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}