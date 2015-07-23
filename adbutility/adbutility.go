package adbutility

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// AdbEndoint interface defines the list of methods that can be called on
// an ADB endpoint. An endpoint can be local or remote based on its
// implementation.
type AdbEndpoint interface {
	Adb(timeout int, args ...string) (string, error)
	GetAttachedDevices(timeout int) ([]string, error)
	WaitForSerials(timeout int, serials ...string) error
	WaitForDevices(timeout int, count int) error
}

// LocalEndpoint struct holds adb executable path for a local adb endpoint and
// implements the interface AdbEndpoint.
type LocalEndpoint struct {
	ADBPath string // Executable path of adb command
}

// GetDefaultEndpoint provides the default adb endpoint implementation, that is,
// local adb where adb command can be located in system path.
func GetDefaultLocalEndpoint() AdbEndpoint {
	return LocalEndpoint{ADBPath: "adb"}
}

// GetLocalEndpoint provides an endpoint on local machine with custome adb
// executable location.
func GetLocalEndpoint(adbPath string) AdbEndpoint {
	return LocalEndpoint{ADBPath: adbPath}
}

// output is a structure that holds adb command's stdin + stdout and error if something
// went wrong while executing the command.
type output struct {
	ret string // Output of a command execution
	err error  // Error of a command execution
}

// outChannel is a channel of output type, and used to get adb command output from
// go routine.
type outChannel chan output

// Execute a ADB command on this local endpoint, it takes integer value timeout,
// which specifies the maximum allowed time to run this command before it get
// killed and a set of arguments as adb command parameters. Please note that it
// panics in case of process can not be killed after timeout. It returns string
// representing the adb command output including stdin and stdout and error is
// returned if something went wrong.
func (ep LocalEndpoint) Adb(timeout int, args ...string) (string, error) {
	log.Println("adb :", args)
	cmd := exec.Command(ep.ADBPath, args...)
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
		if strings.Contains(strings.TrimSpace(out.ret), "device not found") {
			return "", errors.New("Device is disconnected")
		}
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

// GetAttachedDevices method returns list of attached device serial number
// to this local endpoint as a sclice of string. Error is also returned if
// something went wront in adb side.
func (ep LocalEndpoint) GetAttachedDevices(timeout int) ([]string, error) {
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
func (ep LocalEndpoint) WaitForSerials(timeout int, serials ...string) error {
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
func (ep LocalEndpoint) WaitForDevices(timeout int, count int) error {
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
