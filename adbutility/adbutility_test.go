package adbutility

import (
	"testing"
	"time"
)

func TestWaitForDevice(t *testing.T) {
	devs_initial, err := GetAttachedDevices(120)
	if err != nil {
		t.Error(err)
	}

	if len(devs_initial) == 0 {
		err := WaitForDevices(5)
		if err == nil {
			t.Errorf("No initial devices connected, Waiting should have failed")
		}
	} else {
		err := WaitForDevices(5)
		if err != nil {
			t.Errorf("Initial devices were connected, Waiting should have passed")
		}
		err = WaitForDevices(5, devs_initial...)
		if err != nil {
			t.Errorf("Waiting should have passed because all the devices are already connected")
		}
		err = WaitForDevices(5, append(devs_initial, "INVALID_SERIAL_XXYY")...)
		if err == nil {
			t.Errorf("Waiting should have failed because this device should not exist really")
		}
	}

	err = WaitForDevices(10, "INVALID_SERIAL_XXYY")
	if err == nil {
		t.Errorf("Waiting should have failed because this device should not exist really")
	}
}

func TestGetAttachedDevices(t *testing.T) {
	devs_initial, err := GetAttachedDevices(5)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 5)

	devs_final, err := GetAttachedDevices(5)
	if err != nil {
		t.Error(err)
	}

	if len(devs_initial) != len(devs_final) {
		t.Errorf("Number of devices do not match : [%v], [%v]", devs_initial, devs_final)
	}

	for _, dev := range devs_final {
		if !stringInSlice(dev, devs_initial) {
			t.Errorf("Devices do not match : [%v], [%v]", devs_initial, devs_final)
		}
	}
}

func TestAdb(t *testing.T) {
	_, err := Adb(5, "-s", "ASCDEFGHIJKL", "wait-for-device")
	if err == nil {
		t.Errorf("Adb should have failed while waiting for invalid device due to timeout")
	}
	_, err = Adb(5, "ASCDEFGHIJKL")
	if err == nil {
		t.Errorf("Adb should have failed because of unknown command")
	}
}
