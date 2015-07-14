package adbutility

import (
	"fmt"
	"os"
	"testing"
)

func TestLocalEndpoint(t *testing.T) {
	os.Setenv("DEBUG_LOG", "VV")
	ep := GetDefaultLocalEndpoint()
	devs, err := ep.GetAttachedDevices(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println("Devices present :", devs)
	err = ep.WaitForDevices(10, len(devs)+2)
	if err == nil {
		t.Error("It should have failed with timeout")
		t.FailNow()
	}
}
