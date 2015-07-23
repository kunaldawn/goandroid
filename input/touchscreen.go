package input

import (
	"errors"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
	"strconv"
	"strings"
)

// TouchScreen struct represensts touch input susbystem for associated device.
type TouchScreen struct {
	dev  device.Device   // Associated device
	disp display.Display // Associated device display
}

// NewTouchScreen method returns a new TouchScreen and associates it with
// given device.
func NewTouchScreen(dev device.Device) TouchScreen {
	disp := display.NewDisplay(dev)
	return TouchScreen{dev: dev, disp: disp}
}

// Tap method performs a touch operation on specified (x,y) coordinate. It
// returns error on adb operation failure.
func (ts TouchScreen) Tap(x int, y int) error {
	_, err := ts.dev.Shell("input", "tap", strconv.Itoa(x), strconv.Itoa(y))
	return err
}

// Swipe method performs touch swipe operation from given (x1, y1) coordinate
// to (x2, y2) coordinate with specified delay. It returns error on adb operation
// failure.
func (ts TouchScreen) Swipe(x1 int, y1 int, x2 int, y2 int, delay int) error {
	_, err := ts.dev.Shell("input", "touchscreen", "swipe", strconv.Itoa(x1), strconv.Itoa(y1), strconv.Itoa(x2), strconv.Itoa(y2), strconv.Itoa(delay))
	return err
}

// SwipeDown method performs touch swipe down (top --> bottom) operation for
// a number of times defined by given count parameter. It returns error on adb operation failure.
func (ts TouchScreen) SwipeDown(count int) error {
	w, h, err := ts.disp.GetDisplaySize()
	if err != nil {
		return err
	}
	x1 := w / 2
	x2 := x1
	y1 := h / 4
	y2 := y1 * 3
	for i := 0; i < count; i++ {
		err := ts.Swipe(x1, y1, x2, y2, 1000)
		if err != nil {
			return err
		}
	}
	return nil
}

// SwipeUp method performs touch swipe up (bottom --> top) operation for
// a number of times defined by given count parameter. It returns error on
// adb operation failure.
func (ts TouchScreen) SwipeUp(count int) error {
	w, h, err := ts.disp.GetDisplaySize()
	if err != nil {
		return err
	}
	x1 := w / 2
	x2 := x1
	y2 := h / 4
	y1 := y2 * 3
	for i := 0; i < count; i++ {
		err := ts.Swipe(x1, y1, x2, y2, 1000)
		if err != nil {
			return err
		}
	}
	return nil
}

// SwipeLeft method performs touch swipe left (right --> left) operation for
// a number of times defined by given count parameter. It returns error on
// adb operation failure.
func (ts TouchScreen) SwipeLeft(count int) error {
	w, h, err := ts.disp.GetDisplaySize()
	if err != nil {
		return err
	}
	x2 := w / 4
	x1 := x2 * 3
	y2 := h / 2
	y1 := y2
	for i := 0; i < count; i++ {
		err := ts.Swipe(x1, y1, x2, y2, 1000)
		if err != nil {
			return err
		}
	}
	return nil
}

// SwipeRight method performs touch swipe right (right --> left) operation for
// a number of times defined by given count parameter. It returns error on
// adb operation failure.
func (ts TouchScreen) SwipeRight(count int) error {
	w, h, err := ts.disp.GetDisplaySize()
	if err != nil {
		return err
	}
	x1 := w / 4
	x2 := x1 * 3
	y2 := h / 2
	y1 := y2
	for i := 0; i < count; i++ {
		err := ts.Swipe(x1, y1, x2, y2, 1000)
		if err != nil {
			return err
		}
	}
	return nil
}

// RawSendEvent sends raw touch input event on given touch device, it takes
// event type, event code and event value as parameter and returns error on
// adb operation failure. make sure you are using correct device path for
// touch device, and can be obtailed easily by GetTouchInputDevice method.
func (ts TouchScreen) RawSendEvent(dev string, eventType int, event int, value int) error {
	_, err := ts.dev.Shell("sendevent", dev, strconv.Itoa(eventType), strconv.Itoa(event), strconv.Itoa(value))
	return err
}

// GetTouchInputDevice method is used to determine correct touch input device
// path on associated android device. It returns error on adb operation failure or
// if device path can not be determined for any reason.
func (ts TouchScreen) GetTouchInputDevice() (string, error) {
	tag1 := "KEY (0001):"
	tag2 := "ABS (0003):"
	out, err := ts.dev.Shell("getevent", "-p")
	if err != nil {
		return "", err
	}
	lines := strings.Split(out, "\n")

	currentDevice := ""
	tag1_match := false
	tag2_match := false
	for _, line := range lines {

		if strings.Contains(line, "add device") {
			tag1_match = false
			tag2_match = false
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return "", errors.New("Unable to parse device information")
			}
			currentDevice = strings.TrimSpace(parts[1])
			continue
		}

		if strings.Contains(line, tag1) {
			tag1_match = true
			continue
		}

		if strings.Contains(line, tag2) {
			tag2_match = true
			continue
		}

		if tag1_match && tag2_match {
			return currentDevice, nil
		}
	}
	return "", errors.New("Unable to determine touch device")
}
