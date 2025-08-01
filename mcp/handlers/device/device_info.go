package device

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DeviceInfoHandler handles comprehensive device information operations
type DeviceInfoHandler struct {
	manager goandroid.AndroidManager
}

// NewDeviceInfoHandler creates a new device info handler
func NewDeviceInfoHandler(manager goandroid.AndroidManager) *DeviceInfoHandler {
	return &DeviceInfoHandler{
		manager: manager,
	}
}

// RegisterTools registers all device information tools
func (h *DeviceInfoHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Get comprehensive device information
	getDeviceSpecsTool := mcpLib.NewTool("get_device_specs",
		mcpLib.WithDescription("Get comprehensive device specifications including model, OS version, screen size, etc."),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getDeviceSpecsTool, h.handleGetDeviceSpecs)

	// Check device root status
	checkRootStatusTool := mcpLib.NewTool("check_root_status",
		mcpLib.WithDescription("Check if the device has root access"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(checkRootStatusTool, h.handleCheckRootStatus)

	// Get device uptime
	getDeviceUptimeTool := mcpLib.NewTool("get_device_uptime",
		mcpLib.WithDescription("Get device uptime information"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getDeviceUptimeTool, h.handleGetDeviceUptime)

	// Get device battery information
	getBatteryInfoTool := mcpLib.NewTool("get_battery_info",
		mcpLib.WithDescription("Get device battery status and information"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getBatteryInfoTool, h.handleGetBatteryInfo)

	// Get device memory information
	getMemoryInfoTool := mcpLib.NewTool("get_memory_info",
		mcpLib.WithDescription("Get device memory usage information"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getMemoryInfoTool, h.handleGetMemoryInfo)

	// Get device storage information
	getStorageInfoTool := mcpLib.NewTool("get_storage_info",
		mcpLib.WithDescription("Get device storage usage information"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("The serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getStorageInfoTool, h.handleGetStorageInfo)
}

func (h *DeviceInfoHandler) handleGetDeviceSpecs(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	// Gather comprehensive device information
	specs := make(map[string]interface{})

	// Basic device info
	if model, err := device.Device.GetProperty("ro.product.model"); err == nil {
		specs["model"] = model
	}
	if brand, err := device.Device.GetProperty("ro.product.brand"); err == nil {
		specs["brand"] = brand
	}
	if manufacturer, err := device.Device.GetProperty("ro.product.manufacturer"); err == nil {
		specs["manufacturer"] = manufacturer
	}

	// OS Information
	if version, err := device.Device.GetProperty("ro.build.version.release"); err == nil {
		specs["android_version"] = version
	}
	if sdk, err := device.Device.GetProperty("ro.build.version.sdk"); err == nil {
		if sdkInt, err := strconv.Atoi(sdk); err == nil {
			specs["sdk_version"] = sdkInt
		}
	}
	if buildId, err := device.Device.GetProperty("ro.build.id"); err == nil {
		specs["build_id"] = buildId
	}

	// Hardware info
	if cpu, err := device.Device.GetProperty("ro.product.cpu.abi"); err == nil {
		specs["cpu_architecture"] = cpu
	}

	// Screen information
	if width, height, err := device.Display.GetDisplaySize(); err == nil {
		specs["screen_size"] = map[string]interface{}{
			"width":  width,
			"height": height,
		}
	}

	// Network info
	if wifi, err := device.Device.Shell("dumpsys", "wifi"); err == nil {
		specs["wifi_available"] = strings.Contains(wifi, "Wi-Fi is enabled")
	}

	specs["serial"] = args.Serial

	jsonBytes, _ := json.Marshal(specs)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *DeviceInfoHandler) handleCheckRootStatus(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	// Check if device has root access
	rootOutput, err := device.Device.Shell("id")
	isRooted := err == nil && strings.Contains(rootOutput, "uid=0(root)")

	// Also check for su binary
	suOutput, suErr := device.Device.Shell("which", "su")
	hasSuBinary := suErr == nil && strings.TrimSpace(suOutput) != ""

	data := map[string]interface{}{
		"serial":        args.Serial,
		"is_rooted":     isRooted,
		"has_su_binary": hasSuBinary,
		"root_output":   rootOutput,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *DeviceInfoHandler) handleGetDeviceUptime(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	uptimeOutput, err := device.Device.Shell("uptime")
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get uptime: %v", err)), nil
	}

	data := map[string]interface{}{
		"serial": args.Serial,
		"uptime": strings.TrimSpace(uptimeOutput),
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *DeviceInfoHandler) handleGetBatteryInfo(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	batteryOutput, err := device.Device.Shell("dumpsys", "battery")
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get battery info: %v", err)), nil
	}

	// Parse battery information
	batteryInfo := make(map[string]interface{})
	lines := strings.Split(batteryOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "level:") {
			if parts := strings.Split(line, ":"); len(parts) == 2 {
				if level, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					batteryInfo["level"] = level
				}
			}
		} else if strings.Contains(line, "status:") {
			if parts := strings.Split(line, ":"); len(parts) == 2 {
				batteryInfo["status"] = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "health:") {
			if parts := strings.Split(line, ":"); len(parts) == 2 {
				batteryInfo["health"] = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "temperature:") {
			if parts := strings.Split(line, ":"); len(parts) == 2 {
				if temp, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					batteryInfo["temperature"] = temp / 10.0 // Convert to Celsius
				}
			}
		}
	}

	data := map[string]interface{}{
		"serial":       args.Serial,
		"battery_info": batteryInfo,
		"raw_output":   batteryOutput,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *DeviceInfoHandler) handleGetMemoryInfo(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	memoryOutput, err := device.Device.Shell("cat", "/proc/meminfo")
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get memory info: %v", err)), nil
	}

	data := map[string]interface{}{
		"serial":         args.Serial,
		"memory_details": memoryOutput,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}

func (h *DeviceInfoHandler) handleGetStorageInfo(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
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

	storageOutput, err := device.Device.Shell("df", "-h")
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get storage info: %v", err)), nil
	}

	data := map[string]interface{}{
		"serial":          args.Serial,
		"storage_details": storageOutput,
	}

	jsonBytes, _ := json.Marshal(data)
	return mcpLib.NewToolResultText(string(jsonBytes)), nil
}