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

func (views Views) GetByTextIndex(text string, index int) (View, bool) {
	idx := 0
	for _, vw := range views {
		if vw.Text == text {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByMatchingTextIndex(text string, index int) (View, bool) {
	idx := 0
	for _, vw := range views {
		if strings.Contains(strings.ToLower(vw.Text), strings.ToLower(text)) {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}
