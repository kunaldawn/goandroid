// TODO : Documentation

package goandroid

import (
	"github.com/kunaldawn/goandroid/adbutility"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/view"
)

type Android struct {
	Device device.Device
	Input  input.InputManager
	View   view.DeviceView
	// TODO : Define following api
	// Activity   interface{}
	// Package    interface{}
	// Telephoney interface{}
	// Settings   interface{}
}

func NewAndroidDevice(serial string, timeout int) Android {
	dev := device.NewDevice(serial, timeout)
	inp := input.NewInputManager(dev)
	viw := view.NewDeviceView(dev)
	return Android{dev, inp, viw}
}

func GetAttachedAndroidDevices(timeout int) ([]Android, error) {
	serials, err := adbutility.GetAttachedDevices(timeout)
	if err != nil {
		return []Android{}, err
	}
	devices := []Android{}
	for index := range serials {
		dev := NewAndroidDevice(serials[index], timeout)
		devices = append(devices, dev)
	}
	return devices, nil
}
