// TODO : Documentation

package input

import (
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
)

type InputManager struct {
	dev         device.Device
	TouchScreen TouchScreen
	Key         Key
}

func NewTouchScreen(dev device.Device) TouchScreen {
	disp := display.NewDisplay(dev)
	return TouchScreen{dev: dev, disp: disp}
}

func NewInputManager(dev device.Device) InputManager {
	ts := NewTouchScreen(dev)
	key := NewKey(dev)
	return InputManager{dev: dev, TouchScreen: ts, Key: key}
}
