package common

import (
	"encoding/json"

	mcpLib "github.com/mark3labs/mcp-go/mcp"
)

// NewSuccessResult creates a successful tool result with message
func NewSuccessResult(message string, data map[string]interface{}) *mcpLib.CallToolResult {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["success"] = true
	data["message"] = message
	
	return mcpLib.NewToolResultText(toJSON(data))
}

// NewErrorResult creates an error tool result
func NewErrorResult(message string) *mcpLib.CallToolResult {
	return mcpLib.NewToolResultError(message)
}

// NewDataResult creates a tool result with data
func NewDataResult(data map[string]interface{}) *mcpLib.CallToolResult {
	return mcpLib.NewToolResultText(toJSON(data))
}

// toJSON converts data to JSON string
func toJSON(data interface{}) string {
	// Simple JSON conversion - in production you'd want proper error handling
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}