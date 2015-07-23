package input

import (
	"github.com/kunaldawn/goandroid/geometry"
	"time"
)

const (
	EV_ABS                  = 3   // ABS Event
	EV_SYN                  = 0   // Sync Event
	EV_KEY                  = 1   // Key event
	BTN_TOUCH               = 330 // Touch event
	BTN_TOOL_FINGER         = 325 // Finger event
	DOWN                    = 1   // Touch down event
	UP                      = 0   // Touch up event
	ABS_MT_TRACKING_ID      = 57  // ID of the touch (important for multi-touch reports)
	ABS_MT_TOUCH_MAJOR      = 48  // Touch size in pixels
	ABS_MT_POSITION_X       = 53  // X coordinate of the touch
	ABS_X                   = 0   // X coordinate of touch in emulator
	ABS_MT_POSITION_Y       = 54  // Y coordinate of the touch
	ABS_Y                   = 1   // Y coordinate of touch in emulator
	ABS_MT_PRESSURE         = 58  // Pressure of the touch
	SYN_MT_REPORT           = 2   // End of separate touch data
	SYN_REPORT              = 0   // End of report
	DEFAULT_TOUCH_ID        = 0   // Default touch point id
	DEFAULT_PRESSURE        = 50  // Touch pressure default value
	DEFAULT_FINGER_TIP_SIZE = 5   // Default touch finger tip size
)

// DrawGesture method allows you to draw a continious gesture on device view.
// It takes a set of points and a delay parameter, gesture is drawn based on
// given set of points and each point is iterated after specific delay. Please
// note that the delay here is in milliseconds. It returns error on adb
// operation errors. Gesture is drawn based on multitouch protocol v2 events.
func (ts TouchScreen) DrawGesture(points geometry.Points, delay int) error {
	dev, err := ts.GetTouchInputDevice()
	if err != nil {
		return err
	}
	for index, pt := range points {
		if index == 0 {
			err = ts.RawSendEvent(dev, EV_KEY, BTN_TOUCH, DOWN)
			if err != nil {
				return err
			}
		}
		err = ts.RawMovePoint(dev, pt.X, pt.Y, DEFAULT_TOUCH_ID, DEFAULT_PRESSURE, DEFAULT_FINGER_TIP_SIZE)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	err = ts.RawSendEvent(dev, EV_ABS, ABS_MT_TRACKING_ID, 4294967295)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_KEY, BTN_TOUCH, UP)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_KEY, BTN_TOOL_FINGER, UP)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}

// DrawGestureEmulator method allows you to draw gesture on emulator device.
// Emulator devices are generally single touch and operates on different
// touch protocol than of multitouch v2 protocol. This method takes a set of
// points and a delay parameter, gesture is drawn based on given set of points
// and iterated one by one with specific delay in between. Please note tha
// value of delay is in milliseconds. It returns error on adb operation failure.
func (ts TouchScreen) DrawGestureEmulator(points geometry.Points, delay int) error {
	dev, err := ts.GetTouchInputDevice()
	if err != nil {
		return err
	}
	for index, pt := range points {
		err = ts.RawMovePointEmulator(dev, pt.X, pt.Y)
		if err != nil {
			return err
		}
		if index == 0 {
			err = ts.RawSendEvent(dev, EV_KEY, BTN_TOUCH, DOWN)
			if err != nil {
				return err
			}
			err = ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	err = ts.RawSendEvent(dev, EV_KEY, BTN_TOUCH, UP)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}

// RawMovePoint method allows you to move a gesture point on device screen.
// This can be used to draw gesture with different parameters than of defauls.
// Please note that users need to send sync signals by them self, check
// DrawGesture method implementation regarding this, if user specific gesture
// method method implementation is required. This method takes x and y coordinates
// of new gesture point along with touch point id, touch pressure and touch
// finger tip size and touch device path. It returns error on adb operation failure.
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
	err = ts.RawSendEvent(dev, EV_ABS, ABS_MT_PRESSURE, pressure)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}

// RawMovePointEmulator method allows you to move a gesture point on emulator screen.
// This can be used to draw gesture with different parameters than of defauls.
// Please note that users need to send sync signals by them self, check
// DrawGestureEmulator method implementation regarding this, if user specific gesture
// method method implementation is required. This method takes x and y coordinates
// of new gesture point along with touch device path and returns error on adb operation failure.
func (ts TouchScreen) RawMovePointEmulator(dev string, x int, y int) error {
	err := ts.RawSendEvent(dev, EV_ABS, ABS_X, x)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
	if err != nil {
		return err
	}
	err = ts.RawSendEvent(dev, EV_ABS, ABS_Y, y)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}
