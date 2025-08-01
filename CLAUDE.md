# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is **goandroid**, an Android automation library written in Go that uses ADB (Android Debug Bridge) to interact with Android devices. The library provides a pure Go implementation without requiring Java backends, UI Automator wrappers, or installing APKs on devices.

## Architecture

The library follows a modular architecture with these core components:

### Core Library
- **AndroidManager**: Factory for creating Android device instances with configurable ADB timeout
- **Android**: Main device interface that aggregates all subsystems
- **Device**: Low-level ADB communication layer (`device/device.go`)
- **Input**: Touch, key, and text input management (`input/input.go`, `input/touchscreen.go`, `input/key.go`, `input/text.go`)
- **View**: UI element discovery and interaction via uiautomator XML dumps (`view/view.go`)
- **Activity**: Android activity lifecycle management (`activity/activity.go`)
- **Display**: Screen information and display properties (`display/display.go`)
- **Geometry**: Coordinate and rectangle utilities (`geometry/geometry.go`)
- **AdbUtility**: Core ADB command execution (`adbutility/adbutility.go`)

### MCP Server Architecture
The MCP (Model Context Protocol) server exposes all goandroid functionality through 134 tools across 27 handler files:

- **Server**: Dual transport support (stdio/HTTP) with StreamableHTTPServer (`mcp/server.go`)
- **Registry**: Central handler registration system (`mcp/registry/registry.go`)
- **Handlers**: Organized by functionality (`mcp/handlers/`) with categories:
  - Device management (6 handlers): Basic operations, file management, properties, ADB utilities, device info
  - Input operations (6 handlers): Basic input, swipes, gestures, shortcuts, advanced touch
  - View operations (14 handlers): Finding, matching, clicking, scrolling, presence checking, cross-references
  - Activity/Display/Geometry handlers (3 handlers): Lifecycle, screen properties, coordinate utilities

## Common Commands

### Building and Testing
```bash
go build ./...                              # Build all packages
go test ./...                               # Run tests (if available)
go mod tidy                                 # Clean up dependencies
cd cmd/mcp-server && go build -o goandroid-mcp-server .  # Build MCP server
```

### MCP Server Development
```bash
# Build and run MCP server in stdio mode
cd cmd/mcp-server && go run . --adb-path="adb" --timeout=60

# Build and run MCP server in HTTP mode  
cd cmd/mcp-server && go run . --adb-path="adb" --timeout=60 --http=":8080"

# Build executable
cd cmd/mcp-server && go build -o goandroid-mcp-server .
```

### Development
```bash
go run example.go       # Run example code
go doc ./...            # Generate documentation  
gofmt -w .             # Format code
```

## Key Implementation Details

- All operations are ADB-based and require `adb` executable in system PATH
- Device interactions use XML hierarchy dumps from `uiautomator dump`
- The View struct represents UI elements with calculated center points for easy clicking
- Input operations support touchscreen taps, key events, and text input
- Activity management includes launching apps and waiting for focus
- Root access is required for some operations (handled automatically)

## Dependencies

- **External**: ADB (Android Debug Bridge) executable must be available in system PATH
- **Internal**: Uses only Go standard library and internal packages

## Usage Pattern

1. Create AndroidManager with timeout and ADB path
2. Get Android device instance by serial number
3. Use subsystem interfaces (Input, View, Activity, etc.) for automation
4. All operations return errors for proper error handling

## Module Structure

- `go.mod`: Go 1.24.4, module path `github.com/kunaldawn/goandroid`
- Main package exports AndroidManager and Android types
- Subpackages are organized by functionality (activity, input, view, etc.)
- MCP server provides comprehensive API coverage through Protocol tools

## MCP Development Notes

- Handler pattern used for organizing MCP tools by functionality
- Import cycle resolution through `mcp/common/helpers.go` package
- All handlers follow consistent error handling and JSON marshaling patterns
- Device instances validated using struct zero value comparison: `if device == (goandroid.Android{}) {}`
- Tool registration uses `mcpServer.AddTool()` pattern with handler methods