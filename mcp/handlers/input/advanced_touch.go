package input

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AdvancedTouchHandler handles advanced touch operations
type AdvancedTouchHandler struct {
	manager goandroid.AndroidManager
}

// NewAdvancedTouchHandler creates a new advanced touch handler
func NewAdvancedTouchHandler(manager goandroid.AndroidManager) *AdvancedTouchHandler {
	return &AdvancedTouchHandler{manager: manager}
}

// RegisterTools registers advanced touch operation tools
func (h *AdvancedTouchHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get touch input device
	getTouchDeviceTool := mcpLib.NewTool("get_touch_input_device",
		mcpLib.WithDescription("Get the touch input device path for raw touch events"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getTouchDeviceTool, h.handleGetTouchInputDevice)

	// Send raw touch event
	rawSendEventTool := mcpLib.NewTool("raw_send_event",
		mcpLib.WithDescription("Send raw touch event to the device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("device",
			mcpLib.Required(),
			mcpLib.Description("Touch device path (use get_touch_input_device to get)"),
		),
		mcpLib.WithNumber("event_type",
			mcpLib.Required(),
			mcpLib.Description("Event type code"),
		),
		mcpLib.WithNumber("event",
			mcpLib.Required(),
			mcpLib.Description("Event code"),
		),
		mcpLib.WithNumber("value",
			mcpLib.Required(),
			mcpLib.Description("Event value"),
		),
	)
	mcpServer.AddTool(rawSendEventTool, h.handleRawSendEvent)

	// Raw move point
	rawMovePointTool := mcpLib.NewTool("raw_move_point",
		mcpLib.WithDescription("Move a gesture point with raw touch events"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("device",
			mcpLib.Required(),
			mcpLib.Description("Touch device path"),
		),
		mcpLib.WithNumber("x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate"),
		),
		mcpLib.WithNumber("y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate"),
		),
		mcpLib.WithNumber("id",
			mcpLib.Required(),
			mcpLib.Description("Touch point ID"),
		),
		mcpLib.WithNumber("pressure",
			mcpLib.Required(),
			mcpLib.Description("Touch pressure"),
		),
		mcpLib.WithNumber("size",
			mcpLib.Required(),
			mcpLib.Description("Touch size"),
		),
	)
	mcpServer.AddTool(rawMovePointTool, h.handleRawMovePoint)

	// Raw move point for emulator
	rawMovePointEmulatorTool := mcpLib.NewTool("raw_move_point_emulator",
		mcpLib.WithDescription("Move a gesture point on emulator devices"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("device",
			mcpLib.Required(),
			mcpLib.Description("Touch device path"),
		),
		mcpLib.WithNumber("x",
			mcpLib.Required(),
			mcpLib.Description("X coordinate"),
		),
		mcpLib.WithNumber("y",
			mcpLib.Required(),
			mcpLib.Description("Y coordinate"),
		),
	)
	mcpServer.AddTool(rawMovePointEmulatorTool, h.handleRawMovePointEmulator)
}

func (h *AdvancedTouchHandler) handleGetTouchInputDevice(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string `json:"serial"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	devicePath, err := device.Input.TouchScreen.GetTouchInputDevice()
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"serial":      args.Serial,
		"device_path": devicePath,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdvancedTouchHandler) handleRawSendEvent(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial    string  `json:"serial"`
		Device    string  `json:"device"`
		EventType float64 `json:"event_type"`
		Event     float64 `json:"event"`
		Value     float64 `json:"value"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	err := device.Input.TouchScreen.RawSendEvent(
		args.Device,
		int(args.EventType),
		int(args.Event),
		int(args.Value),
	)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Raw touch event sent",
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdvancedTouchHandler) handleRawMovePoint(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial   string  `json:"serial"`
		Device   string  `json:"device"`
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		ID       float64 `json:"id"`
		Pressure float64 `json:"pressure"`
		Size     float64 `json:"size"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	err := device.Input.TouchScreen.RawMovePoint(
		args.Device,
		int(args.X),
		int(args.Y),
		int(args.ID),
		int(args.Pressure),
		int(args.Size),
	)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Raw move point executed",
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *AdvancedTouchHandler) handleRawMovePointEmulator(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string  `json:"serial"`
		Device string  `json:"device"`
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	err := device.Input.TouchScreen.RawMovePointEmulator(
		args.Device,
		int(args.X),
		int(args.Y),
	)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Raw move point (emulator) executed",
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}