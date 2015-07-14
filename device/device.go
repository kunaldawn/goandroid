// Package device provides utility methods that provides interfaces to invoke adb
// commands and adb shell commands on
package device

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/adbutility"
	"strings"
	"time"
)

const (
	PROP_BOOT_STATUS = "sys.boot_completed" // Property that represents devices boot status
)

type Device struct {
	Serial      string                 // Device serial number
	Timeout     int                    // Timeout in seconds for all adb and shell operations
	AdbEndpoint adbutility.AdbEndpoint // Adb endpoint for this device
}

// NewDevice creates a new Device based on given serial number, adb operation timeout and
// connecting adb enpoint for this device.
func NewDevice(serial string, timeout int, endPoint adbutility.AdbEndpoint) Device {
	return Device{Serial: serial, Timeout: timeout, AdbEndpoint: endPoint}
}

// IsAvailable checkes availability of the device in the adb endpoint. It returns
// boolean value indicating availibility status of the device and error in case of
// something went wrong while doing adb operations.
func (dev Device) IsAvailable() (bool, error) {
	devices, err := dev.AdbEndpoint.GetAttachedDevices(dev.Timeout)
	if err != nil {
		return false, err
	}
	for index := range devices {
		if dev.Serial == devices[index] {
			return true, nil
		}
	}
	return false, nil
}

// Adb allows to execute adb command on this device instance. It takes a adb command
// and optional list of arguments. It returns outout of the adb command and error in case
// of something went wrong. Please note that adb will timeout within default adb operation
// timeout.
func (dev Device) Adb(command string, args ...string) (string, error) {
	return dev.AdbEndpoint.Adb(dev.Timeout, append([]string{"-s", dev.Serial, command}, args...)...)
}

func (dev Device) Shell(command string, args ...string) (string, error) {
	return dev.Adb("shell", append([]string{command}, args...)...)
}

func (dev Device) GetProperty(key string) (string, error) {
	prop, err := dev.GetAllProperties()
	if err != nil {
		return "", err
	}
	val, ok := prop[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("Key [%s] is not found in device properties", key))
	}
	return val, nil
}

func (dev Device) GetAllProperties() (map[string]string, error) {
	prop_map := make(map[string]string)
	prop, err := dev.Shell("getprop")
	if err != nil {
		return prop_map, err
	}
	lines := strings.Split(prop, "\n")
	for index := range lines {
		parts := strings.Split(lines[index], ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(strings.Replace(strings.Replace(parts[0], "[", "", -1), "]", "", -1))
			value := strings.TrimSpace(strings.Replace(strings.Replace(parts[1], "[", "", -1), "]", "", -1))
			prop_map[key] = value
		}
	}
	return prop_map, nil
}

func (dev Device) Pull(src string, dst string) (string, error) {
	return dev.Adb("pull", src, dst)
}

func (dev Device) Push(src string, dst string) (string, error) {
	return dev.Adb("push", src, dst)
}

func (dev Device) WaitForAvailability(timeout int) error {
	_, err := dev.AdbEndpoint.Adb(timeout, "-s", dev.Serial, "wait-for-device")
	return err
}

func (dev Device) Root() error {
	out, err := dev.Adb("root")
	if err != nil {
		return err
	}
	if !strings.Contains(out, "restarting adbd as root") && !strings.Contains(out, "adbd is already running as root") {
		return errors.New("Unable to gain root access to device")
	}
	return dev.WaitForAvailability(dev.Timeout)
}

func (dev Device) Reboot(restartTimeout int, bootTimeout int) error {
	_, err := dev.Adb("reboot")
	if err != nil {
		return err
	}
	err = dev.WaitForAvailability(restartTimeout)
	if err != nil {
		return err
	}
	return dev.WaitForBootToComplete(bootTimeout)
}

func (dev Device) WaitForBootToComplete(bootTimeout int) error {
	startTime := time.Now()
	for {
		currentTime := time.Now()
		delta := currentTime.Sub(startTime)
		if delta.Seconds() >= float64(bootTimeout) {
			break
		}
		val, err := dev.GetProperty(PROP_BOOT_STATUS)
		if err != nil {
			return err
		}
		if val == "1" {
			return nil
		}
	}
	return errors.New("Device not completed boot sequence in timeout")
}
