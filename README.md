![goandroid_logo](goandroid_logo.png)

goandroid
=========
[![Project Status](http://stillmaintained.com/kunaldawn/goandroid.png)](https://stillmaintained.com/kunaldawn/goandroid) [![Build Status](http://img.shields.io/travis/kunaldawn/goandroid.svg?style=flat-square)](https://travis-ci.org/kunaldawn/goandroid) [![Coverage Status](http://img.shields.io/coveralls/kunaldawn/goandroid.svg?style=flat-square)](https://coveralls.io/r/kunaldawn/goandroid) [![Issues](http://img.shields.io/github/issues/kunaldawn/goandroid.svg?style=flat-square)](https://github.com/kunaldawn/goandroid/issues) [![Documentation](http://img.shields.io/badge/go-Documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/kunaldawn/goandroid) [![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/kunaldawn/goandroid/blob/master/LICENSE)

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