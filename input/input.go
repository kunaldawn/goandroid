package input

import (
	"github.com/kunaldawn/goandroid/device"
)

// InputManager struct holds all type of supported device input interfaces
// such as touch screen input interface to operate touch based inputes, key
// input interface to operate on key press inputs, text input interface to
// insert text on selected text box (text input function is very limited at this
// point, does not supports unicode and any special characters and its limited
// by "adb shell input text capabilites").
type InputManager struct {
	dev         device.Device // Associated device
	TouchScreen TouchScreen   // Associated touch screen input
	Key         Key           // Associated key input
	TextInput   TextInput     // Associated text input
}

// NewInputManager method returns a new InputManager struct and associates
// it with given device.
func NewInputManager(dev device.Device) InputManager {
	ts := NewTouchScreen(dev)
	key := NewKey(dev)
	ti := NewTextInput(dev)
	return InputManager{dev: dev, TouchScreen: ts, Key: key, TextInput: ti}
}
