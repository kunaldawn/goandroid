package input

import (
	"github.com/kunaldawn/goandroid/geometry"
	"time"
)

func (ts TouchScreen) DrawGestureProtocolV1(points geometry.Points, delay int) error {
	dev, err := ts.GetTouchInputDevice()
	if err != nil {
		return err
	}
	for _, pt := range points {
		// Move the point
		err = ts.RawMovePointProtocolV1(dev, pt.X, pt.Y, DEFAULT_TOUCH_ID, DEFAULT_PRESSURE, DEFAULT_FINGER_TIP_SIZE)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
		// Send SYNC for this movement
		err = ts.RawSendEvent(dev, EV_SYN, SYN_MT_REPORT, 0)
		if err != nil {
			return err
		}
		err = ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
		if err != nil {
			return err
		}
	}
	// Send SYNC to release the point
	err = ts.RawSendEvent(dev, EV_SYN, SYN_MT_REPORT, 0)
	if err != nil {
		return err
	}
	return ts.RawSendEvent(dev, EV_SYN, SYN_REPORT, 0)
}

func (ts TouchScreen) DrawGestureProtocolV2(points geometry.Points, delay int) error {
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
		err = ts.RawMovePointProtocolV2(dev, pt.X, pt.Y, DEFAULT_TOUCH_ID, DEFAULT_PRESSURE, DEFAULT_FINGER_TIP_SIZE)
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

func (ts TouchScreen) RawMovePointProtocolV1(dev string, x int, y int, id int, pressure int, size int) error {
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

func (ts TouchScreen) RawMovePointProtocolV2(dev string, x int, y int, id int, pressure int, size int) error {
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
