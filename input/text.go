package input

import (
	"github.com/kunaldawn/goandroid/device"
	"strings"
)

// TextInput struct represents a text input subsystem for associated device.
type TextInput struct {
	dev device.Device // Associated device
}

// NewTextInput method returns a new TextInput struct which is associated with
// given device.
func NewTextInput(dev device.Device) TextInput {
	return TextInput{dev: dev}
}

// EnterText method enters text on selected input area. Input area must be
// selected previously. Functionality of this method is very limited and
// does not support any unicode aharacters or any special characters.
func (ti TextInput) EnterText(text string) {
	formatted := strings.Replace(text, " ", "%s", -1)
	ti.dev.Shell("input", "text", formatted)
}
