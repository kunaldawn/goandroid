package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DeviceHandler handles device management tools
type DeviceHandler struct {
	manager goandroid.AndroidManager
}

// NewDeviceHandler creates a new device handler
func NewDeviceHandler(manager goandroid.AndroidManager) *DeviceHandler {
	return &DeviceHandler{manager: manager}
}

// RegisterTools registers device management tools
func (h *DeviceHandler) RegisterTools(mcpServer *server.MCPServer) {
	// List attached devices tool
	listDevicesTool := mcpLib.NewTool("list_devices",
		mcpLib.WithDescription("List all attached Android devices"),
	)
	mcpServer.AddTool(listDevicesTool, h.handleListDevices)

	// Get device info tool
	getDeviceInfoTool := mcpLib.NewTool("get_device_info",
		mcpLib.WithDescription("Get information about a specific Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
	)
	mcpServer.AddTool(getDeviceInfoTool, h.handleGetDeviceInfo)

	// Shell command tool
	shellCommandTool := mcpLib.NewTool("shell_command",
		mcpLib.WithDescription("Execute a shell command on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("command",
			mcpLib.Required(),
			mcpLib.Description("Shell command to execute"),
		),
		mcpLib.WithArray("args",
			mcpLib.Description("Command arguments"),
		),
	)
	mcpServer.AddTool(shellCommandTool, h.handleShellCommand)

	// Install APK tool
	installApkTool := mcpLib.NewTool("install_apk",
		mcpLib.WithDescription("Install an APK file on the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("apk_path",
			mcpLib.Required(),
			mcpLib.Description("Path to the APK file to install"),
		),
	)
	mcpServer.AddTool(installApkTool, h.handleInstallApk)

	// Uninstall app tool
	uninstallAppTool := mcpLib.NewTool("uninstall_app",
		mcpLib.WithDescription("Uninstall an application from the Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("package_name",
			mcpLib.Required(),
			mcpLib.Description("Package name of the application to uninstall"),
		),
	)
	mcpServer.AddTool(uninstallAppTool, h.handleUninstallApp)
}

// handleListDevices handles the list_devices tool
func (h *DeviceHandler) handleListDevices(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	devices, err := h.manager.GetAttachedAndroidDevices()
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get attached devices: %v", err)), nil
	}

	deviceList := make([]string, len(devices))
	for i, device := range devices {
		deviceList[i] = device.Device.Serial
	}

	result := map[string]interface{}{
		"devices": deviceList,
		"count":   len(deviceList),
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleGetDeviceInfo handles the get_device_info tool
func (h *DeviceHandler) handleGetDeviceInfo(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	
	// Get device properties
	props, err := device.Device.Shell("getprop")
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to get device properties: %v", err)), nil
	}

	result := map[string]interface{}{
		"serial":     serial,
		"properties": props,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleShellCommand handles the shell_command tool
func (h *DeviceHandler) handleShellCommand(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	command, err := request.RequireString("command")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	args := request.GetArguments()
	var cmdArgs []string
	if argsVal, ok := args["args"].([]interface{}); ok {
		for _, arg := range argsVal {
			if argStr, ok := arg.(string); ok {
				cmdArgs = append(cmdArgs, argStr)
			}
		}
	}

	device := h.manager.GetNewAndroidDevice(serial)
	allArgs := append([]string{command}, cmdArgs...)
	var output string
	var shellErr error
	if len(allArgs) == 1 {
		output, shellErr = device.Device.Shell(allArgs[0])
	} else {
		output, shellErr = device.Device.Shell(allArgs[0], allArgs[1:]...)
	}
	if shellErr != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to execute shell command: %v", shellErr)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"output":  output,
		"command": command,
		"args":    cmdArgs,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleInstallApk handles the install_apk tool
func (h *DeviceHandler) handleInstallApk(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	apkPath, err := request.RequireString("apk_path")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	output, err := device.Device.Adb("install", apkPath)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to install APK: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("APK installed successfully: %s", apkPath),
		"output":  output,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handleUninstallApp handles the uninstall_app tool
func (h *DeviceHandler) handleUninstallApp(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	packageName, err := request.RequireString("package_name")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	output, err := device.Device.Adb("uninstall", packageName)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to uninstall app: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("App uninstalled successfully: %s", packageName),
		"output":  output,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}