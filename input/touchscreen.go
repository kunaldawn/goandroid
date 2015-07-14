package input

import (
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
	"strconv"
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
