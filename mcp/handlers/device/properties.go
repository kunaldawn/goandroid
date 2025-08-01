package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// PropertiesHandler handles device property operations
type PropertiesHandler struct {
	manager goandroid.AndroidManager
}

// NewPropertiesHandler creates a new properties handler
func NewPropertiesHandler(manager goandroid.AndroidManager) *PropertiesHandler {
	return &PropertiesHandler{
		manager: manager,
	}
}

// RegisterTools registers all property-related tools
func (h *PropertiesHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get specific device property
	getPropertyTool := mcpLib.NewTool("get_device_property",
		mcpLib.WithDescription("Get a specific property from an Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithString("key",
			mcpLib.Required(),
			mcpLib.Description("The property key to retrieve (e.g., 'ro.build.version.release')"),
		),
	)
	mcpServer.AddTool(getPropertyTool, h.handleGetDeviceProperty)

	// Get all device properties
	getAllPropertiesTool := mcpLib.NewTool("get_all_device_properties",
		mcpLib.WithDescription("Get all properties from an Android device as a map"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getAllPropertiesTool, h.handleGetAllDeviceProperties)
}

func (h *PropertiesHandler) handleGetDeviceProperty(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial string `json:"serial"`
		Key    string `json:"key"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	value, err := device.Device.GetProperty(args.Key)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"serial": args.Serial,
		"key":    args.Key,
		"value":  value,
	}
	
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *PropertiesHandler) handleGetAllDeviceProperties(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	properties, err := device.Device.GetAllProperties()
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"serial":     args.Serial,
		"properties": properties,
	}
	
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}