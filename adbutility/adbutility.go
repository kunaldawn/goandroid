// TODO : Documentation

package adbutility

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/logging"
	"os/exec"
	"strings"
	"time"
)

type output struct {
	ret string // Output of a command execution
	err error  // Error of a command execution
}

type outChannel chan output

func Adb(timeout int, args ...string) (string, error) {
	logging.LogVV("Adb : args [%v]", args)
	cmd := exec.Command("adb", args...)
	done := make(outChannel)
	go func(done outChannel, cmd *exec.Cmd) {
		// Run the command and get combined output for stdout and stderr
		ret, err := cmd.CombinedOutput()
		// Create a output for command
		out := output{ret: string(ret), err: err}
		// Notify that command has complited
		done <- out
	}(done, cmd)
	select {
	case out := <-done:
		return out.ret, out.err
	case <-time.After(time.Duration(timeout) * time.Second):
		err := cmd.Process.Kill()
		if err != nil {
			panic(err)
		}
		return "", errors.New(fmt.Sprintf("Timeout occured after %d sec while executing adb %v", timeout, args))
	}
	return "", nil
}

func GetAttachedDevices(timeout int) ([]string, error) {
	logging.LogVV("GetAttachedDevices : timeout [%d]", timeout)
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

func WaitForDevices(timeout int, serials ...string) error {
	logging.LogVV("WaitForDevices : timeout [%d] : serials [%v]", timeout, serials)
	innter_to := 5
	if timeout < innter_to {
		innter_to = timeout
	}
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		devs, err := GetAttachedDevices(innter_to)
		if err != nil {
			return err
		}
		if len(serials) == 0 {
			if len(devs) > 0 {
				return nil
			}
		} else {
			found := 0
			for _, dev := range devs {
				if stringInSlice(dev, serials) {
					found += 1
				}
			}
			if found == len(serials) {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d sec while waiting for devices", timeout))
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
