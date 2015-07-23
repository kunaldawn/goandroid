package adbutility

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// AdbEndpoint struct defines an adb communication endpoint for an android device.
type AdbEndpoint struct {
	ADBPath string // Executable path of adb command
}

// GetNewAdbEndpoint method returns a new ADBEndpoint instance with specified adb
// executable path. If adb tool is on your sustem path, then just pass "adb", else
// mention the full path for your adb executable.
func GetNewAdbEndpoint(adb string) AdbEndpoint {
	return AdbEndpoint{ADBPath: adb}
}

// AdbEndpointOutput is a struct that holds adb command's stdin + stdout and error if something
// went wrong while executing the command.
type AdbEndpointOutput struct {
	Ret string // Output of an adb command execution
	Err error  // Error of an adb command execution
}

// AdbEndpointOutputChannel is a channel of AdbEndpointOutput type, and used to get adb command output from
// go routine.
type AdbEndpointOutputChannel chan AdbEndpointOutput

// Execute an ADB command on this endpoint, it takes integer value timeout,
// which specifies the maximum allowed time to run this command before it gets
// killed and a set of arguments as adb command parameters. Please note that it
// panics in case of process can not be killed after timeout. It returns string
// representing the adb command output including stdout and stderr and error is
// returned if something went wrong.
func (ep AdbEndpoint) Adb(timeout int, args ...string) (string, error) {
	log.Println("adb :", args)
	cmd := exec.Command(ep.ADBPath, args...)
	done := make(AdbEndpointOutputChannel)
	go func(done AdbEndpointOutputChannel, cmd *exec.Cmd) {
		// Run the command and get combined output for stdout and stderr
		ret, err := cmd.CombinedOutput()
		// Create a output for command
		out := AdbEndpointOutput{Ret: string(ret), Err: err}
		// Notify that command has complited
		done <- out
	}(done, cmd)
	select {
	case out := <-done:
		if strings.Contains(strings.TrimSpace(out.Ret), "device not found") {
			return "", errors.New("Device is disconnected")
		}
		return out.Ret, out.Err
	case <-time.After(time.Duration(timeout) * time.Second):
		err := cmd.Process.Kill()
		if err != nil {
			panic(err)
		}
		return "", errors.New(fmt.Sprintf("Timeout occured after %d sec while executing adb %v", timeout, args))
	}
	return "", nil
}

// GetAttachedDevices method returns list of attached device serial number
// to this local endpoint as a sclice of string. Error is also returned if
// something went wront in adb side.
func (ep AdbEndpoint) GetAttachedDevices(timeout int) ([]string, error) {
	devices := []string{}
	adb_output, err := ep.Adb(timeout, "devices")
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

// WaitForSerial waits for speified set of serial numbers to be available on this
// local endpoint. It returns error in case of all specified serial numbers can not
// be detected by this endpoint. Error is also returned if something went wront in
// adb side.
func (ep AdbEndpoint) WaitForSerials(timeout int, serials ...string) error {
	if len(serials) == 0 {
		return errors.New(fmt.Sprintf("No serials specified", timeout))
	}

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
		devs, err := ep.GetAttachedDevices(innter_to)
		if err != nil {
			return err
		}
		found := 0
		for _, dev := range devs {
			for _, item := range serials {
				if dev == item {
					found += 1
					break
				}
			}
		}
		if found == len(serials) {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d sec while waiting for serials", timeout))
}

// WaitForDevices method waits for specified number of devices to be available to this
// local endpoint. If specific number of devices can not be located in timeout, it returns
// error. Error is also returned if something went wront in adb side.
func (ep AdbEndpoint) WaitForDevices(timeout int, count int) error {
	if count <= 0 {
		return errors.New(fmt.Sprintf("Can not wait for less than or eual to zero devices", timeout))
	}

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
		devs, err := ep.GetAttachedDevices(innter_to)
		if err != nil {
			return err
		}
		if count == len(devs) {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d sec while waiting for devices", timeout))
}
