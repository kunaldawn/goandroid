package input

import (
	"errors"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
	"github.com/kunaldawn/goandroid/geometry"
	"strconv"
	"strings"
	"time"
)

const (
	EV_ABS                  = 3  // ABS Event
	EV_SYN                  = 0  // Sync Event
	ABS_MT_TRACKING_ID      = 57 // ID of the touch (important for multi-touch reports)
	ABS_MT_TOUCH_MAJOR      = 48 // Touch size in pixels
	ABS_MT_POSITION_X       = 53 // X coordinate of the touch
	ABS_MT_POSITION_Y       = 54 // Y coordinate of the touch
	ABS_MT_PRESSURE         = 58 // Pressure of the touch
	SYN_MT_REPORT           = 2  // End of separate touch data
	SYN_REPORT              = 0  // End of report
	DEFAULT_TOUCH_ID        = 50 // Default touch point id
	DEFAULT_PRESSURE        = 50 // Touch pressure default value
	DEFAULT_FINGER_TIP_SIZE = 5  // Default touch finger tip size
)

type TouchScreen struct {
	dev  device.Device
	disp display.Display
}

func (ts TouchScreen) Tap(x int, y int) error {
	_, err := ts.dev.Shell("input", "tap", strconv.Itoa(x), strconv.Itoa(y))
	return err
}

func (ts TouchScreen) Swipe(x1 int, y1 int, x2 int, y2 int, delay int) error {
	_, err := ts.dev.Shell("input", "touchscreen", "swipe", strconv.Itoa(x1), strconv.Itoa(y1), strconv.Itoa(x2), strconv.Itoa(y2), strconv.Itoa(delay))
	return err
}

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

func (ts TouchScreen) DrawGesture(points geometry.Points, delay int) error {
	dev, err := ts.GetTouchInputDevice()
	if err != nil {
		return err
	}
	for _, pt := range points {
		err = ts.RawMovePoint(dev, pt.X, pt.Y, DEFAULT_TOUCH_ID, DEFAULT_PRESSURE, DEFAULT_FINGER_TIP_SIZE)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
		err = ts.RawEndReport(dev)
		if err != nil {
			return err
		}
	}
	return ts.RawEndReport(dev)
}

func (ts TouchScreen) GetTouchInputDevice() (string, error) {
	tag1 := "KEY (0001): 0213"
	tag2 := "ABS (0003): 0035"
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

func (ts TouchScreen) RawSendEvent(dev string, eventType int, event int, value int) error {
	_, err := ts.dev.Shell("sendevent", dev, strconv.Itoa(eventType), strconv.Itoa(event), strconv.Itoa(value))
	return err
}

func (ts TouchScreen) RawMovePoint(dev string, x int, y int, id int, pressure int, size int) error {
	err := ts.RawSendEvent(dev, EV_ABS, ABS_MT_TRACKING_ID, id)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_ABS, ABS_MT_POSITION_X, x)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_ABS, ABS_MT_POSITION_Y, y)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_ABS, ABS_MT_TOUCH_MAJOR, size)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_ABS, ABS_MT_PRESSURE, pressure)
}

func (ts TouchScreen) RawEndReport(dev string) error {
	err := ts.RawSendEvent(dev, EV_SYN, SYN_MT_REPORT, 0)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}
