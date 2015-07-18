// TODO : Documentation

package goandroid

import (
	"github.com/kunaldawn/goandroid/activity"
	"github.com/kunaldawn/goandroid/adbutility"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/view"
)

type AndroidManager struct {
	endpoint adbutility.AdbEndpoint
	timeout  int
}

type Android struct {
	Device   device.Device
	Input    input.InputManager
	View     view.DeviceView
	Display  display.Display
	Activity activity.Activity
}

func GetDefaultAndroidManager(timeout int) AndroidManager {
	return AndroidManager{endpoint: adbutility.GetDefaultLocalEndpoint(), timeout: timeout}
}

func (am AndroidManager) NewAndroidDevice(serial string) Android {
	dev := device.NewDevice(serial, am.timeout, am.endpoint)
	inp := input.NewInputManager(dev)
	viw := view.NewDeviceView(dev)
	disp := display.NewDisplay(dev)
	act := activity.NewActivity(dev)
	return Android{dev, inp, viw, disp, act}
}

func (am AndroidManager) GetAttachedAndroidDevices(timeout int) ([]Android, error) {
	serials, err := am.endpoint.GetAttachedDevices(timeout)
	if err != nil {
		return []Android{}, err
	}
	devices := []Android{}
	for index := range serials {
		dev := am.NewAndroidDevice(serials[index])
		devices = append(devices, dev)
	}
	return devices, nil
}
