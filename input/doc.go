// Package input provides a set of useful methods that helps on input based interaction
// with an associated android device. All input interactions are either based on "adb shell input"
// command or via sending raw input commands to input device directly.
//
// Following are the list of supported input mechanisms :
//    * Touch screen based input
//    * Key press based input
//    * Simple text input
//
// Touch screen based input operations provide following features:
//    * Perform single touch on any screen coordinate
//    * Draw multi point gesture on real device or emulator
//    * Perform swipe operation between any two points in the screen
//    * Perform generic swipe operations such as swipe up/down/left/right on screen
//    * Determine devices touch input device path
//    * Send raw touch input commands to touch input device
//
// Key press based input operations provide following features:
//    * Send any key press code for any number of times
//    * Provides generic key press methods such as home/power/back...
//
// Text input operations provide following features:
//    * Allows you to enter simple text on devices focused input area
//
// Currently text input operation is very limited and can not insert any
// special characters or unicode characters. This limitation is imposed by
// "adb shell input text" command.
package input
