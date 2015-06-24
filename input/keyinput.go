package input

import (
	"github.com/kunaldawn/goandroid/device"
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

type Key struct {
	dev device.Device
}

func NewKey(dev device.Device) Key {
	return Key{dev: dev}
}

func (key Key) Press(code int) error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(code))
	return err
}

func (key Key) PressMenu() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(MENU_KEY))
	return err
}

func (key Key) PressHome() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(HOME_KEY))
	return err
}

func (key Key) PressBack() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(BACK_KEY))
	return err
}

func (key Key) PressCall() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(CALL_KEY))
	return err
}

func (key Key) PressEndCall() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(ENDCALL_KEY))
	return err
}

func (key Key) PressUp() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(UP_KEY))
	return err
}

func (key Key) PressDown() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(DOWN_KEY))
	return err
}

func (key Key) PressLeft() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(LEFT_KEY))
	return err
}

func (key Key) PressRight() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(RIGHT_KEY))
	return err
}

func (key Key) PressVolumeUp() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(VOLUME_UP_KEY))
	return err
}

func (key Key) PressVolumeDown() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(VOLUME_DOWN_KEY))
	return err
}

func (key Key) PressPower() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(POWER_KEY))
	return err
}

func (key Key) PressCamera() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(CAMERA_KEY))
	return err
}

func (key Key) PressEnter() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(ENTER_KEY))
	return err
}

func (key Key) PressDelete() error {
	_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(DEL_KEY))
	return err
}
