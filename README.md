![goandroid_logo](goandroid_logo.png)

goandroid

Introduction
------------
**"goandroid" is an Android automation library purely written in GO.**

**Project Status :** Under Development, not yet ready for v1.0 release.

Whether you are an android developer and want to do some automation tasks on your android device to reduce some manual human work, or an enthusiastic developer who want to do some automation taks on your android device, this library allows you to write automation code and do cool stuffs on your android device.

#### TODO's
- Complete all Package Documentation's
- Implement all features for v1.0
- Write test code for all packages such that it can be tested on Travis CI over emulator

#### FAQ's
- **Is it a wrapper arround Android [UI Automator](https://developer.android.com/tools/testing-support-library/index.html#UIAutomator) library?**

    No, goandroid does not uses UI Automator for its implementation. It does not uses any java backend or JSON RPC service to communicate UI Automator.

- **Is it Android testing framework?**  

    No, goandroid is not an android trsting framework. Its an automation framework, but can be used for android UI automation and testing perpuses also.

- **Does it installs anything on my android device?**

    No, goandroid does not installs any APK or other tools in your android device to provide any features.

- **Can you explain how does it work?**

    This library is purely based on ADB (Android Debugger Bridge) and adb tool. It uses adb commands to perform all operations on device. You can write automation code using this library and check logs for what adb command has been executed for that automation logic.

- **I want this feature X, can you include X in goandroid?**

    If the feature you are requesting can be implemented only by using adb commands, yes I will add the feature for you. Just make a pull request or start a new issue describing the adb equivalent implementation for the feature.

- **What are the dependencies of goandroid?**

    The only dependency of goandroid library is "adb" executable tool.

MCP Integration & Claude Code
------------------------------

**goandroid** now includes comprehensive **Model Context Protocol (MCP)** support, enabling seamless integration with Claude Code and other MCP-compatible AI development tools.

### What is MCP Integration?

The MCP server exposes all goandroid functionality through **134 standardized tools** across **27 handler categories**, allowing AI assistants like Claude Code to directly control Android devices through natural language commands.

### Integration with Claude Code

Claude Code can now perform Android automation tasks by communicating with the goandroid MCP server:

```bash
# Start MCP server in stdio mode (for Claude Code)
cd cmd/mcp-server && go run . --adb="adb" --timeout=60

# Or start in HTTP mode
cd cmd/mcp-server && go run . --adb="adb" --timeout=60 --http=":8080"
```

#### Registering HTTP MCP Server with Claude Code CLI

Once your HTTP MCP server is running, register it with Claude Code CLI:

```bash
# Register the HTTP MCP server (replace with your server URL and port)
claude mcp add --transport http goandroid-server http://localhost:8080

# List registered MCP servers to verify
claude mcp list

# Get details about the registered server
claude mcp get goandroid-server
```

**Configuration Options:**
- Use `--scope local` for project-specific configuration (default)
- Use `--scope user` to make the server available across all projects
- Add authentication headers if needed: `--header "Authorization: Bearer your-token"`

**Server Management:**
```bash
# Remove a registered server
claude mcp remove goandroid-server

# Update server configuration
claude mcp remove goandroid-server
claude mcp add --transport http goandroid-server http://localhost:8080
```

### Available MCP Tools (134 Total)

#### Device Management Tools
- **Device Operations**: Check availability, get properties, execute ADB commands
- **File Management**: Push/pull files, manage device storage
- **Device Information**: Battery status, memory info, system specs, root detection
- **Device Control**: Reboot, wait for boot completion, root access

#### Input Automation Tools  
- **Basic Input**: Tap, swipe, enter text
- **Advanced Touch**: Multi-touch gestures, raw touch events
- **Key Events**: Hardware keys, navigation, volume controls
- **Gesture Drawing**: Complex gesture patterns, emulator-specific gestures

#### UI Interaction Tools
- **Element Finding**: Locate UI elements by text, resource ID, description, type
- **Presence Checking**: Verify element existence and visibility
- **Click Operations**: Click elements by various selectors
- **Scrolling**: Scroll to find elements, directional scrolling
- **State Management**: Check element states, get element properties
- **Cross-Reference**: Get properties of related elements

#### Activity & Display Tools
- **Activity Management**: Start apps, check focused activity, wait for focus
- **Display Control**: Get/set screen dimensions, display properties
- **Geometry Utilities**: Calculate distances, rectangle operations, coordinate math

### Example: Using Claude Code with goandroid MCP

Once the MCP server is running, Claude Code can execute commands like:

```
"Tap on the Settings button"
"Scroll down to find Developer Options and enable it"
"Check if the device is rooted"
"Get the current battery level"
"Take a screenshot and analyze the UI"
```

The MCP server automatically translates these requests into appropriate ADB commands and device interactions.

### Benefits of MCP Integration

- **Natural Language Control**: Control Android devices through conversational AI
- **Comprehensive Coverage**: All goandroid APIs exposed as MCP tools
- **Dual Transport**: Support for both stdio and HTTP protocols
- **Error Handling**: Robust error reporting and validation
- **Development Acceleration**: Rapid prototyping of Android automation workflows

Usage & Example
----------------

### Install adb
First of all make sure you have "adb" tool in your system path.

For Ubuntu 14.04 or later use following commands to install adb:
```bash
sudo apt-get update
sudo apt-get install android-tools-adb
```

For other distributions, download Android SDK and "adb" tool can be located inside"platform-tools" direcory. Now add this to your system path using following comands:
```bash
cd <root folder of sdk>
export PATH=$PATH:$PWD/platform-tools
```

### Initialize Android device instance
First import "github.com/kunaldawn/goandroid" in your source, and you are ready to write some cool automation code. To interact with an android device, you need to create an android device instance first. Following example shows how to create a new android device instance and enable "Show CPU Usage" settings using automation. Please locate the documentation for package [goandroid](https://godoc.org/github.com/kunaldawn/goandroid) for more information.

**Example:**

**[Youtube Screen Cast](https://www.youtube.com/watch?v=vuq2Cq82xJ4)**

```go
package main

import (
	"github.com/kunaldawn/goandroid"
)

func main() {
	// Creat a new android manager with 60 seconds adb time out and take adb
	// executable path from system path.
	android_manager := goandroid.GetNewAndroidManager(60, "adb")

	// Create an android device instance with following serial
	android := android_manager.GetNewAndroidDevice("emulator-5554")

	// Start settings activity
	android.Activity.StartActivity("com.android.settings")
	// Wait for settings activity to get focused and displayed on screen
	// with 10 seconds timeout
	android.Activity.WaitForActivityToFocus("com.android.settings", 10)

	// Scroll down to "About phone"
	android.View.ScrollDownToText("About phone", 0, 10)
	// Now click "About phone" settings item
	android.View.ClickText("About phone", 0, 5)

	// Now scroll down to "Build number"
	android.View.ScrollDownToText("Build number", 0, 10)

	// Now for faster click operation, we are going to get the view for "Build number" text
	view, _ := android.View.GetViewForText("Build number", 0, 5)

	// Now we will click the text 10 times
	for i := 0; i < 10; i++ {
		android.Input.TouchScreen.Tap(view.Center.X, view.Center.Y)
	}

	// Now go back to main settings page
	android.Input.Key.PressBack(1)
	// Click developer options
	android.View.ClickText("Developer options", 0, 10)
	
	// Now scroll down to "Show CPU Usage" and enable it
	android.View.ScrollDownToMatchingText("show cpu", 0, 10)
	android.View.ClickMatchingText("show cpu", 0, 10)
}
```

**Translated adb commands by goandroid for above code:**
```
adb : [-s emulator-5554 root]
adb : [-s emulator-5554 wait-for-device]
adb : [-s emulator-5554 shell am start com.android.settings]
adb : [-s emulator-5554 shell dumpsys activity | grep mFocusedActivity:]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell input tap 369 1643]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input tap 188 1684]
adb : [-s emulator-5554 shell input keyevent 4]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell input tap 433 1426]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell wm size]
adb : [-s emulator-5554 shell input touchscreen swipe 540 1440 540 480 1000]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell uiautomator dump]
adb : [-s emulator-5554 shell cat /storage/sdcard/window_dump.xml]
adb : [-s emulator-5554 shell input tap 229 1677]
```