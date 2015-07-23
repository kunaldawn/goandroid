package input

import (
	"github.com/kunaldawn/goandroid/device"
	"strconv"
)

const (
	MENU_KEY        = 1  // Menu key code
	HOME_KEY        = 3  // Home key code
	BACK_KEY        = 4  // Back key code
	CALL_KEY        = 5  // Cll key code
	ENDCALL_KEY     = 6  // End call key code
	UP_KEY          = 19 // Up key code
	DOWN_KEY        = 20 // Down key code
	LEFT_KEY        = 21 // Left key code
	RIGHT_KEY       = 22 // Right key code
	VOLUME_UP_KEY   = 24 // Volume up key code
	VOLUME_DOWN_KEY = 25 // Volume down key code
	POWER_KEY       = 26 // Power key code
	CAMERA_KEY      = 27 // Camera key code
	ENTER_KEY       = 66 // Enter key code
	DEL_KEY         = 67 // Del key code
)

// Key struct defines a key input subsystem associated with a device.
type Key struct {
	dev device.Device // Associated device
}

// NewKey method returns a new Key struct associated with given device.
func NewKey(dev device.Device) Key {
	return Key{dev: dev}
}

// Press method performs a key press event based on given key code, this operation
// is repeated based on the count parameter. It returns error on adb operation
// failure.
func (key Key) Press(code int, count int) error {
	for i := 0; i < count; i++ {
		_, err := key.dev.Shell("input", "keyevent", strconv.Itoa(code))
		if err != nil {
			return err
		}
	}
	return nil
}

// PressMenu method performs a menu key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressMenu(count int) error {
	return key.Press(MENU_KEY, count)
}

// PressHome method performs a home key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressHome(count int) error {
	return key.Press(HOME_KEY, count)
}

// PressBack method performs a back key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressBack(count int) error {
	return key.Press(BACK_KEY, count)
}

// PressCall method performs a call key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressCall(count int) error {
	return key.Press(CALL_KEY, count)
}

// PressEndCall method performs a end call key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressEndCall(count int) error {
	return key.Press(ENDCALL_KEY, count)
}

// PressUp method performs a up key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressUp(count int) error {
	return key.Press(UP_KEY, count)
}

// PressDown method performs a down key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressDown(count int) error {
	return key.Press(DOWN_KEY, count)
}

// PressLeft method performs a left key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressLeft(count int) error {
	return key.Press(LEFT_KEY, count)
}

// PressRight method performs a right key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressRight(count int) error {
	return key.Press(RIGHT_KEY, count)
}

// PressVolumeUp method performs a volume up key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressVolumeUp(count int) error {
	return key.Press(VOLUME_UP_KEY, count)
}

// PressVolumeDown method performs a volume down key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressVolumeDown(count int) error {
	return key.Press(VOLUME_DOWN_KEY, count)
}

// PressPower method performs a power key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressPower(count int) error {
	return key.Press(POWER_KEY, count)
}

// PressCamera method performs a camera key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressCamera(count int) error {
	return key.Press(CAMERA_KEY, count)
}

// PressEnter method performs a enter key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressEnter(count int) error {
	return key.Press(ENTER_KEY, count)
}

// PressDelete method performs a delete key press event, this operation is repeated
// based on the count parameter. It returns error on adb operation failure.
func (key Key) PressDelete(count int) error {
	return key.Press(DEL_KEY, count)
}
