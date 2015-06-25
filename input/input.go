// TODO : Documentation

package input

import (
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/logging"
)

type InputManager struct {
	dev         device.Device
	TouchScreen TouchScreen
	Key         Key
}

func NewTouchScreen(dev device.Device) TouchScreen {
	logging.LogVV("NewTouchScreen : device [%v]", dev)
	return TouchScreen{dev: dev}
}

func NewInputManager(dev device.Device) InputManager {
	logging.LogVV("NewInputManager : device [%v]", dev)
	ts := NewTouchScreen(dev)
	key := NewKey(dev)
	return InputManager{dev: dev, TouchScreen: ts, Key: key}
}
