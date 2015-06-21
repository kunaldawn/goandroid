// TODO : Documentation

package input

import (
	"github.com/kunaldawn/go-android/device"
	"strconv"
)

const (
	MENU_KEY        = 1
	HOME_KEY        = 3
	BACK_KEY        = 4
	CALL_KEY        = 5
	ENDCALL_KEY     = 6
	UP_KEY          = 19
	DOWN_KEY        = 20
	LEFT_KEY        = 21
	RIGHT_KEY       = 22
	VOLUME_UP_KEY   = 24
	VOLUME_DOWN_KEY = 25
	POWER_KEY       = 26
	CAMERA_KEY      = 27
	ENTER_KEY       = 66
	DEL_KEY         = 67
)

type TouchScreen struct {
	dev device.Device
}

type Key struct {
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

func NewKey(dev device.Device) Key {
	return Key{dev: dev}
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

func (key Key) Press(code int) error {
	_, err := ts.dev.Shell("input", "keyevent", strconv.Itoa(code))
	return err
}

// TODO : Determine what public api is required for Key

// TODO : Determine what public api is required for touchscreen, such as drag, swipe longpress etc etc
