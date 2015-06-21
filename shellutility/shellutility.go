// TODO : Documentation

package shellutility

import (
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type output struct {
	ret string // Output of a command execution
	err error  // Error of a command execution
}

type outChannel chan output

func RunShellCommand(timeout int, command string, args ...string) (string, error) {
	fmt.Println(fmt.Sprintf("%s %v", command, args))
	cmd := exec.Command(command, args...)
	done := make(outChannel)
	go runShellCommand(done, cmd)
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

func runShellCommand(done outChannel, cmd *exec.Cmd) {
	// Run the command and get combined output for stdout and stderr
	ret, err := cmd.CombinedOutput()
	// Create a output for command
	out := output{ret: string(ret), err: err}
	// Notify that command has complited
	done <- out
}

// TODO : Determine what other public API endpoints are required based on geeral use cases
