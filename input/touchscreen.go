package input

import (
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/logging"
	"strconv"
)

type TouchScreen struct {
	dev device.Device
}

func (ts TouchScreen) Tap(x int, y int) error {
	logging.Log("Tap : x [%d] : y [%d]", x, y)
	_, err := ts.dev.Shell("input", "tap", strconv.Itoa(x), strconv.Itoa(y))
	return err
}
