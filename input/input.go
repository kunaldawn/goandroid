// TODO : Documentation

package input

import (
	"github.com/kunaldawn/goandroid/device"
	"strconv"
)

type TouchScreen struct {
	dev device.Device
}

type InputManager struct {
	dev         device.Device
	TouchScreen TouchScreen
	Key         Key
}

func NewTouchScreen(dev device.Device) TouchScreen {
	return TouchScreen{dev: dev}
}

func NewInputManager(dev device.Device) InputManager {
	ts := NewTouchScreen(dev)
	key := NewKey(dev)
	return InputManager{dev: dev, TouchScreen: ts, Key: key}
}

func (ts TouchScreen) Tap(x int, y int) error {
	_, err := ts.dev.Shell("input", "tap", strconv.Itoa(x), strconv.Itoa(y))
	return err
}

// TODO : Determine what public api is required for Key

// TODO : Determine what public api is required for touchscreen, such as drag, swipe longpress etc etc
