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

func (key Key) Press(code int, count int) error {
	for i := 0; i < count; i++ {
		_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(code))
		if err != nil {
			return err
		}
	}
	return nil
}

func (key Key) PressMenu(count int) error {
	return key.Press(MENU_KEY, count)
}

func (key Key) PressHome(count int) error {
	return key.Press(HOME_KEY, count)
}

func (key Key) PressBack(count int) error {
	return key.Press(BACK_KEY, count)
}

func (key Key) PressCall(count int) error {
	return key.Press(CALL_KEY, count)
}

func (key Key) PressEndCall(count int) error {
	return key.Press(ENDCALL_KEY, count)
}

func (key Key) PressUp(count int) error {
	return key.Press(UP_KEY, count)
}

func (key Key) PressDown(count int) error {
	return key.Press(DOWN_KEY, count)
}

func (key Key) PressLeft(count int) error {
	return key.Press(LEFT_KEY, count)
}

func (key Key) PressRight(count int) error {
	return key.Press(RIGHT_KEY, count)
}

func (key Key) PressVolumeUp(count int) error {
	return key.Press(VOLUME_UP_KEY, count)
}

func (key Key) PressVolumeDown(count int) error {
	return key.Press(VOLUME_DOWN_KEY, count)
}

func (key Key) PressPower(count int) error {
	return key.Press(POWER_KEY, count)
}

func (key Key) PressCamera(count int) error {
	return key.Press(CAMERA_KEY, count)
}

func (key Key) PressEnter(count int) error {
	return key.Press(ENTER_KEY, count)
}

func (key Key) PressDelete(count int) error {
	return key.Press(DEL_KEY, count)
}
