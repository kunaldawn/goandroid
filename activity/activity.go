package activity

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/device"
	"strings"
	"time"
)

// Activity struct represents activity subsystem for associated android device.
type Activity struct {
	dev device.Device // Device instance
}

// NewActivity method gives an initialized instance of activity manager.
// It takes device.Device is a parameter and returns an instance of activity.Activity
// struct.
func NewActivity(dev device.Device) Activity {
	return Activity{dev: dev}
}

// StartActivity method launches a activity on device. See "am start" for more
// more details regarding this command. It takes canonilcal class name as its
// first parameter which defines the package and class name of the activity to
// be launched and a list of accepable options by am start command. This method
// requires root access on device and returns error if root access can not be
// granted. It also returns other adb related errors if something goes wrong.
//
// TODO Check if activity can not be launched or not found or other command
// related outputs.
func (am Activity) StartActivity(canonicalClass string, options ...string) error {
	if len(options) > 0 {
		cmd := []string{}
		cmd = append(cmd, "start")
		cmd = append(cmd, options...)
		cmd = append(cmd, canonicalClass)
		_, err := am.dev.Shell("am", cmd...)
		if err != nil {
			return err
		}
	} else {
		_, err := am.dev.Shell("am", "start", canonicalClass)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFocusedActivity method provides currently focused activity on device screen.
// It uses adb command "dumpsys acyivity" command to determine focused activity
// and extracts package name from the line "mFocusedActivity".
// It returens error if something went wrong on adb side or the method is unable to
// determine the activity canonical package name.
func (am Activity) GetFocusedActivity() (string, error) {
	ret, err := am.dev.Shell("dumpsys", "activity", "|", "grep", "mFocusedActivity:")
	if err != nil {
		return "", err
	}
	if !strings.Contains(ret, "mFocusedActivity:") {
		return "", errors.New("Unable to detect focused activity")
	}
	trimmed := strings.TrimSpace(ret)
	parts := strings.Split(trimmed, " ")
	if len(parts) < 3 {
		return "", errors.New("Unable to determine activity canonical package name from line : " + trimmed)
	}
	return parts[3], nil
}

// IsActivityFocused method checks if given activity is focused on device
// screen or not. If the currently focused activity canonical package contains
// given name then it returns true, and its not case sensitive in nature.
// It returns error if activity information can not be parsed or somethis else
// went wront in between.
func (am Activity) IsActivityFocused(name string) (bool, error) {
	activity, err := am.GetFocusedActivity()
	if err != nil {
		return false, err
	}
	if strings.Contains(strings.ToLower(activity), strings.ToLower(name)) {
		return true, nil
	}
	return false, nil
}

// WaitForActivityToFocus method waits for given activity name to be focused on
// device screen in given timeout period. If activity is not focused within timeout
// then error is returned. Also error is returned if something else went wrong such as
// adb error or current focused activity can not be determined. This method returns
// instantly when given activity is found to be focused on screen.
func (am Activity) WaitForActivityToFocus(name string, timeout int) error {
	startTime := time.Now()
	for {
		currentTime := time.Now()
		delta := currentTime.Sub(startTime)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		stat, err := am.IsActivityFocused(name)
		if err != nil {
			return err
		}
		if stat {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Activity %s is not focused on screen within timeout of %d seconds", name, timeout))
}
