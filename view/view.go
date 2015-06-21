// TODO : Documentation
package view

import (
	"strings"
)

type View struct {
	Index         int
	Class         string
	Package       string
	Resource      string
	Text          string
	Description   string
	Clickable     bool
	Checkable     bool
	Checked       bool
	Enabled       bool
	Focusable     bool
	Focused       bool
	Scrollable    bool
	LongClickable bool
	Password      bool
	Selected      bool
	Bound         Rect
	Center        Point
}

type Views []View

func (vw View) Click() error {
	return nil
}

func (views Views) GetViewByText(text string) (View, bool) {
	for _, vw := range views {
		if vw.Text == text {
			return vw, true
		}
	}
	return View{}, false
}

func (views Views) GetViewsByText(text string, index int) (View, bool) {
	for _, vw := range views {
		if vw.Text == text {
			return vw, true
		}
	}
	return View{}, false
}

func (views Views) GetViewByMatchingText(text string) (View, bool) {
	for _, vw := range views {
		if strings.Contains(strings.ToLower(vw.Text), strings.ToLower(text)) {
			return vw, true
		}
	}
	return View{}, false
}
