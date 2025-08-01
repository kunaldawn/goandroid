package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kunaldawn/goandroid"
	mcpLib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// FileHandler handles file transfer tools
type FileHandler struct {
	manager goandroid.AndroidManager
}

// NewFileHandler creates a new file handler
func NewFileHandler(manager goandroid.AndroidManager) *FileHandler {
	return &FileHandler{manager: manager}
}

// RegisterTools registers file transfer tools
func (h *FileHandler) RegisterTools(mcpServer *server.MCPServer) {
	// Push file tool
	pushFileTool := mcpLib.NewTool("push_file",
		mcpLib.WithDescription("Push a file from local system to Android device"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("local_path",
			mcpLib.Required(),
			mcpLib.Description("Local file path to push"),
		),
		mcpLib.WithString("remote_path",
			mcpLib.Required(),
			mcpLib.Description("Remote path on the device"),
		),
	)
	mcpServer.AddTool(pushFileTool, h.handlePushFile)

	// Pull file tool
	pullFileTool := mcpLib.NewTool("pull_file",
		mcpLib.WithDescription("Pull a file from Android device to local system"),
		mcpLib.WithString("serial",
			mcpLib.Required(),
			mcpLib.Description("Serial number of the Android device"),
		),
		mcpLib.WithString("remote_path",
			mcpLib.Required(),
			mcpLib.Description("Remote file path on the device"),
		),
		mcpLib.WithString("local_path",
			mcpLib.Required(),
			mcpLib.Description("Local path to save the file"),
		),
	)
	mcpServer.AddTool(pullFileTool, h.handlePullFile)
}

// handlePushFile handles the push_file tool
func (h *FileHandler) handlePushFile(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	localPath, err := request.RequireString("local_path")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	remotePath, err := request.RequireString("remote_path")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	output, err := device.Device.Push(localPath, remotePath)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to push file: %v", err)), nil
	}

	result := map[string]interface{}{
		"success":     true,
		"message":     fmt.Sprintf("File pushed successfully from %s to %s", localPath, remotePath),
		"local_path":  localPath,
		"remote_path": remotePath,
		"output":      output,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}

// handlePullFile handles the pull_file tool
func (h *FileHandler) handlePullFile(ctx context.Context, request mcpLib.CallToolRequest) (*mcpLib.CallToolResult, error) {
	serial, err := request.RequireString("serial")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	remotePath, err := request.RequireString("remote_path")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	localPath, err := request.RequireString("local_path")
	if err != nil {
		return mcpLib.NewToolResultError(err.Error()), nil
	}

	device := h.manager.GetNewAndroidDevice(serial)
	output, err := device.Device.Pull(remotePath, localPath)
	if err != nil {
		return mcpLib.NewToolResultError(fmt.Sprintf("Failed to pull file: %v", err)), nil
	}

	result := map[string]interface{}{
		"success":     true,
		"message":     fmt.Sprintf("File pulled successfully from %s to %s", remotePath, localPath),
		"remote_path": remotePath,
		"local_path":  localPath,
		"output":      output,
	}

	resultJSON, _ := json.Marshal(result)
	return mcpLib.NewToolResultText(string(resultJSON)), nil
}