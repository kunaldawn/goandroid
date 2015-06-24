// TODO : Documentation

package adbutility

import (
	"errors"
	"fmt"
	"log"
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
	log.Println(fmt.Sprintf("adb %v", args))
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
		return "", errors.New(fmt.Sprintf("Timeout occured after %d sec", timeout))
	}
	return "", nil
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
