package adbutility

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type AdbEndpoint interface {
	Adb(timeout int, args ...string) (string, error)
	GetAttachedDevices(timeout int) ([]string, error)
	WaitForSerials(timeout int, serials ...string) error
	WaitForDevices(timeout int, count int) error
}

type localEndpoint struct {
	adbPath string // Executable path of adb command
}

type remoteEndpoint struct {
	url string // Host name of adb endpoint
}

func GetDefaultLocalEndpoint() AdbEndpoint {
	return localEndpoint{adbPath: "adb"}
}

func GetLocalEndpoint(adbPath string) AdbEndpoint {
	return GetDefaultLocalEndpoint()
}

func GetRemoteEndpoint(url string) AdbEndpoint {
	return GetDefaultLocalEndpoint()
}

type output struct {
	ret string // Output of a command execution
	err error  // Error of a command execution
}

type outChannel chan output

func (ep localEndpoint) Adb(timeout int, args ...string) (string, error) {
	log.Println("adb :", args)
	cmd := exec.Command(ep.adbPath, args...)
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

func (ep localEndpoint) GetAttachedDevices(timeout int) ([]string, error) {
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

func (ep localEndpoint) WaitForSerials(timeout int, serials ...string) error {
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

func (ep localEndpoint) WaitForDevices(timeout int, count int) error {
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
