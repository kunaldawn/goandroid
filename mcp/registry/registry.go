package registry

import (
	"github.com/kunaldawn/goandroid"
	"github.com/kunaldawn/goandroid/mcp/handlers/activity"
	"github.com/kunaldawn/goandroid/mcp/handlers/device"
	"github.com/kunaldawn/goandroid/mcp/handlers/display"
	"github.com/kunaldawn/goandroid/mcp/handlers/geometry"
	"github.com/kunaldawn/goandroid/mcp/handlers/input"
	"github.com/kunaldawn/goandroid/mcp/handlers/view"
	"github.com/mark3labs/mcp-go/server"
)

// HandlerRegistry manages all MCP tool handlers
type HandlerRegistry struct {
	manager goandroid.AndroidManager
}

// NewHandlerRegistry creates a new handler registry
func NewHandlerRegistry(manager goandroid.AndroidManager) *HandlerRegistry {
	return &HandlerRegistry{manager: manager}
}

// RegisterAllTools registers all MCP tools with the server
func (r *HandlerRegistry) RegisterAllTools(mcpServer *server.MCPServer) {
	// Device management tools
	deviceHandler := device.NewDeviceHandler(r.manager)
	deviceHandler.RegisterTools(mcpServer)

	fileHandler := device.NewFileHandler(r.manager)
	fileHandler.RegisterTools(mcpServer)

	propertiesHandler := device.NewPropertiesHandler(r.manager)
	propertiesHandler.RegisterTools(mcpServer)

	managementHandler := device.NewManagementHandler(r.manager)
	managementHandler.RegisterTools(mcpServer)

	adbUtilityHandler := device.NewAdbUtilityHandler(r.manager)
	adbUtilityHandler.RegisterTools(mcpServer)

	deviceInfoHandler := device.NewDeviceInfoHandler(r.manager)
	deviceInfoHandler.RegisterTools(mcpServer)

	// Geometry utility tools
	geometryUtilsHandler := geometry.NewGeometryUtilsHandler()
	geometryUtilsHandler.RegisterTools(mcpServer)

	// Input operation tools
	basicInputHandler := input.NewBasicInputHandler(r.manager)
	basicInputHandler.RegisterTools(mcpServer)

	swipeHandler := input.NewSwipeHandler(r.manager)
	swipeHandler.RegisterTools(mcpServer)

	gestureHandler := input.NewGestureHandler(r.manager)
	gestureHandler.RegisterTools(mcpServer)

	shortcutsHandler := input.NewShortcutsHandler(r.manager)
	shortcutsHandler.RegisterTools(mcpServer)

	additionalKeysHandler := input.NewAdditionalKeysHandler(r.manager)
	additionalKeysHandler.RegisterTools(mcpServer)

	advancedTouchHandler := input.NewAdvancedTouchHandler(r.manager)
	advancedTouchHandler.RegisterTools(mcpServer)

	// View operation tools
	finderHandler := view.NewFinderHandler(r.manager)
	finderHandler.RegisterTools(mcpServer)

	matcherHandler := view.NewMatcherHandler(r.manager)
	matcherHandler.RegisterTools(mcpServer)

	interactionHandler := view.NewInteractionHandler(r.manager)
	interactionHandler.RegisterTools(mcpServer)

	presenceHandler := view.NewPresenceHandler(r.manager)
	presenceHandler.RegisterTools(mcpServer)

	clickHandler := view.NewClickHandler(r.manager)
	clickHandler.RegisterTools(mcpServer)

	scrollingHandler := view.NewScrollingHandler(r.manager)
	scrollingHandler.RegisterTools(mcpServer)

	stateHandler := view.NewStateHandler(r.manager)
	stateHandler.RegisterTools(mcpServer)

	crossRefHandler := view.NewCrossRefHandler(r.manager)
	crossRefHandler.RegisterTools(mcpServer)

	matchingOperationsHandler := view.NewMatchingOperationsHandler(r.manager)
	matchingOperationsHandler.RegisterTools(mcpServer)

	viewRetrievalHandler := view.NewViewRetrievalHandler(r.manager)
	viewRetrievalHandler.RegisterTools(mcpServer)

	scrollExtensionsHandler := view.NewScrollExtensionsHandler(r.manager)
	scrollExtensionsHandler.RegisterTools(mcpServer)

	crossRefExtensionsHandler := view.NewCrossRefExtensionsHandler(r.manager)
	crossRefExtensionsHandler.RegisterTools(mcpServer)

	// Activity management tools
	activityHandler := activity.NewActivityHandler(r.manager)
	activityHandler.RegisterTools(mcpServer)

	// Display management tools
	displayHandler := display.NewDisplayHandler(r.manager)
	displayHandler.RegisterTools(mcpServer)
}