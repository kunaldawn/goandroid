// TODO : Documentation

package adbutility

import (
	"github.com/kunaldawn/go-android/shellutility"
	"strings"
)

func Adb(timeout int, args ...string) (string, error) {
	return shellutility.RunShellCommand(timeout, "adb", args...)
}

func GetAttachedDevices(timeout int) ([]string, error) {
	devices := []string{}
	adb_output, err := Adb(timeout, "devices")
	if err != nil {
		return devices, err
	}
	list := strings.Split(adb_output, "\n")
	for line := range list {
		if strings.Contains(list[line], "device") && !strings.Contains(list[line], "List of") {
			device_id_line := strings.Split(list[line], "device")
			device_id := strings.TrimSpace(device_id_line[0])
			devices = append(devices, device_id)
		}
	}
	return devices, nil
}
