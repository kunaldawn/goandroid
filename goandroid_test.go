package goandroid

import (
	"github.com/kunaldawn/goandroid/adbutility"
	"testing"
)

func TestGetAttachedAndroidDevices(t *testing.T) {
	err := adbutility.WaitForDevices(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	devs, err := GetAttachedAndroidDevices(10)
	if err != nil {
		t.Error(err)
	}
	
	if len(devs) == 0 {
		t.Errorf("Expected atleast one device to be present. But found zero.")
		t.FailNow()
	}
	
	dev0 := devs[0]
	dev0.Input.Key.PressPower(500)
}
