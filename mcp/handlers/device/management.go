package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ManagementHandler handles device management operations
type ManagementHandler struct {
	manager goandroid.AndroidManager
}

// NewManagementHandler creates a new management handler
func NewManagementHandler(manager goandroid.AndroidManager) *ManagementHandler {
	return &ManagementHandler{
		manager: manager,
	}
}

// RegisterTools registers all management-related tools
func (h *ManagementHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Check device availability
	checkAvailabilityTool := mcpLib.NewTool("check_device_availability",
		mcpLib.WithDescription("Check if an Android device is available"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(checkAvailabilityTool, h.handleCheckDeviceAvailability)

	// Wait for device to become available
	waitForDeviceTool := mcpLib.NewTool("wait_for_device",
		mcpLib.WithDescription("Wait for an Android device to become available"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 30)"),
		),
	)
	mcpServer.AddTool(waitForDeviceTool, h.handleWaitForDevice)

	// Root device
	rootDeviceTool := mcpLib.NewTool("root_device",
		mcpLib.WithDescription("Gain root access on the Android device (requires rooted device)"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(rootDeviceTool, h.handleRootDevice)

	// Reboot device
	rebootDeviceTool := mcpLib.NewTool("reboot_device",
		mcpLib.WithDescription("Reboot the Android device and wait for it to come back online"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("restart_timeout",
			mcpLib.Description("Timeout for restart in seconds (default: 30)"),
		),
		mcpLib.WithNumber("boot_timeout",
			mcpLib.Description("Timeout for boot completion in seconds (default: 120)"),
		),
	)
	mcpServer.AddTool(rebootDeviceTool, h.handleRebootDevice)

	// Wait for boot to complete
	waitForBootTool := mcpLib.NewTool("wait_for_boot_complete",
		mcpLib.WithDescription("Wait for the Android device boot sequence to complete"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
		mcpLib.WithNumber("timeout",
			mcpLib.Description("Timeout in seconds (default: 120)"),
		),
	)
	mcpServer.AddTool(waitForBootTool, h.handleWaitForBootComplete)
}

func (h *ManagementHandler) handleCheckDeviceAvailability(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	available, err := device.Device.IsAvailable()
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"serial":    args.Serial,
		"available": available,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ManagementHandler) handleWaitForDevice(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial  string  `json:"serial"`
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

	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 30
	}

	err := device.Device.WaitForAvailability(timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Device is now available",
		"serial":  args.Serial,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ManagementHandler) handleRootDevice(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	err := device.Device.Root()
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Root access granted",
		"serial":  args.Serial,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ManagementHandler) handleRebootDevice(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial         string  `json:"serial"`
		RestartTimeout float64 `json:"restart_timeout"`
		BootTimeout    float64 `json:"boot_timeout"`
	}

	argBytes, _ := json.Marshal(request.Params.Arguments)
	if err := json.Unmarshal(argBytes, &args); err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Invalid arguments: %v", err)), nil
	}

	device := h.manager.GetNewAndroidDevice(args.Serial)
	if device == (goandroid.Android{}) {
		return mcpLib.NewToolResultError("Device not found"), nil
	}

	restartTimeout := int(args.RestartTimeout)
	if restartTimeout <= 0 {
		restartTimeout = 30
	}
	bootTimeout := int(args.BootTimeout)
	if bootTimeout <= 0 {
		bootTimeout = 120
	}

	err := device.Device.Reboot(restartTimeout, bootTimeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Device rebooted successfully",
		"serial":  args.Serial,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *ManagementHandler) handleWaitForBootComplete(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	var args struct {
		Serial  string  `json:"serial"`
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

	timeout := int(args.Timeout)
	if timeout <= 0 {
		timeout = 120
	}

	err := device.Device.WaitForBootToComplete(timeout)
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	data := map[string]interface{}{
		"success": true,
		"message": "Boot sequence completed",
		"serial":  args.Serial,
	}
	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}