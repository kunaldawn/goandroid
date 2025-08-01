// Package mcp provides Model Context Protocol (MCP) server implementation for
// the goandroid automation library. It exposes goandroid's Android device
// automation capabilities as MCP tools that can be consumed by LLM applications.
//
// The MCP server provides tools for:
//   - Device management (listing devices, getting device information, shell commands, APK management, file transfer)
//   - Input operations (touch, key press, text input, swipe gestures, multi-point gestures)
//   - UI element discovery and interaction (finding views by various criteria, getting hierarchy, clicking elements)
//   - Activity management (launching apps, getting current activity, waiting for activity focus)
//   - Display management (getting and setting display size)
//
// Device Management Tools:
//   - list_devices: List all attached Android devices
//   - get_device_info: Get information about a specific device
//   - shell_command: Execute shell commands on the device
//   - install_apk: Install APK files on the device
//   - uninstall_app: Uninstall applications from the device
//   - push_file: Push files from local system to device
//   - pull_file: Pull files from device to local system
//
// Input Operation Tools:
//   - tap: Tap at specific coordinates
//   - key_press: Press hardware keys
//   - text_input: Input text
//   - swipe: Swipe between two points
//   - swipe_up/swipe_down/swipe_left/swipe_right: Directional swipes
//   - draw_gesture: Draw continuous gestures with multiple points
//
// UI Element Tools:
//   - get_view_hierarchy: Get the complete UI hierarchy
//   - find_view_by_text: Find UI elements by exact text match
//   - find_view_by_resource: Find UI elements by resource ID
//   - find_view_by_description: Find UI elements by content description
//   - find_view_by_type: Find UI elements by class type
//   - find_view_by_matching_*: Case-insensitive versions of the above
//   - click_view: Click on UI elements found by various criteria
//
// Activity Management Tools:
//   - launch_activity: Launch Android activities
//   - get_focused_activity: Get the currently focused activity
//   - wait_for_activity: Wait for a specific activity to become focused
//
// Display Management Tools:
//   - get_display_size: Get current display dimensions
//   - set_display_size: Set display dimensions
//
// Usage:
//   server := mcp.NewGoAndroidMCPServer(timeout, adbPath)
//   err := server.ServeStdio()
package mcp