package goandroid

import (
	"github.com/kunaldawn/goandroid/activity"
	"github.com/kunaldawn/goandroid/adbutility"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/display"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/view"
)

// AndroidManager struct defines a android device manager with an associated
// adb endpoint and adb operation timeout. All devices returned by this android
// manader will be having this adb operation timeout.
type AndroidManager struct {
	Endpoint adbutility.AdbEndpoint // Associated adb endpoint
	Timeout  int                    // Default adb operation timeout in seconds
}

// Android struct defines an android device. Android device has raw device
// communication interface via Device struct, device input interaction interface
// via Input struct, device ui view query interface via View struct, device display
// information interface via Display struct and device application activity
// interface via Activity struct. See their respective documentation for
// list of available mechanisms.
type Android struct {
	Device   device.Device      // Raw adb device communication interface
	Input    input.InputManager // Device input interaction interface
	View     view.DeviceView    // Device UI View query interface
	Display  display.Display    // Device display insterface
	Activity activity.Activity  // Device application activity interface
}

// GetAndroidManager method returns a new AndroidManager instance based on specified
// adb executable path and adb operation timeout. Please note that adb executable must
// be present on the specified path. If adb is on system path then just pass "adb".
// Timeout specified is in seconds and all adb commands will timeout after specified
// seconds.
func GetNewAndroidManager(timeout int, adb string) AndroidManager {
	return AndroidManager{Endpoint: adbutility.GetNewAdbEndpoint(adb), Timeout: timeout}
}

// GetNewAndroidDevice method returns a new Android device instance based on the
// specified serial number.
func (am AndroidManager) GetNewAndroidDevice(serial string) Android {
	dev := device.NewDevice(serial, am.Timeout, am.Endpoint)
	inp := input.NewInputManager(dev)
	viw := view.NewDeviceView(dev)
	disp := display.NewDisplay(dev)
	act := activity.NewActivity(dev)
	return Android{dev, inp, viw, disp, act}
}

// GetAttachedAndroidDevices method returns list of attached android devices
// to the system. It returns error if any error occured wile performing adb operation. 
func (am AndroidManager) GetAttachedAndroidDevices() ([]Android, error) {
	serials, err := am.Endpoint.GetAttachedDevices(am.Timeout)
	if err != nil {
		return []Android{}, err
	}
	devices := []Android{}
	for index := range serials {
		dev := am.GetNewAndroidDevice(serials[index])
		devices = append(devices, dev)
	}
	return devices, nil
}
