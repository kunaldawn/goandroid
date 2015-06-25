// TODO : Documentation
package view

import (
	"github.com/kunaldawn/goandroid/logging"
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

func (views Views) GetByText(text string, index int) (View, bool) {
	logging.LogV("GetByText : text [%s] : index [%d]", text, index)
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

func (views Views) GetByMatchingText(text string, index int) (View, bool) {
	logging.LogV("GetByMatchingText : text [%s] : index [%d]", text, index)
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

func (views Views) GetByResource(resource string, index int) (View, bool) {
	logging.LogV("GetByResource : resource [%s] : index [%d]", resource, index)
	idx := 0
	for _, vw := range views {
		if vw.Resource == resource {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByMatchingResource(resource string, index int) (View, bool) {
	logging.LogV("GetByMatchingResource : resource [%s] : index [%d]", resource, index)
	idx := 0
	for _, vw := range views {
		if strings.Contains(strings.ToLower(vw.Resource), strings.ToLower(resource)) {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByDescription(description string, index int) (View, bool) {
	logging.LogV("GetByDescription : description [%s] : index [%d]", description, index)
	idx := 0
	for _, vw := range views {
		if vw.Description == description {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByMatchingDescription(description string, index int) (View, bool) {
	logging.LogV("GetByMatchingDescription : description [%s] : index [%d]", description, index)
	idx := 0
	for _, vw := range views {
		if strings.Contains(strings.ToLower(vw.Description), strings.ToLower(description)) {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByType(typename string, index int) (View, bool) {
	logging.LogV("GetByType : typename [%s] : index [%d]", typename, index)
	idx := 0
	for _, vw := range views {
		if vw.Class == typename {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}

func (views Views) GetByMatchingType(typename string, index int) (View, bool) {
	logging.LogV("GetByMatchingType : typename [%s] : index [%d]", typename, index)
	idx := 0
	for _, vw := range views {
		if strings.Contains(strings.ToLower(vw.Class), strings.ToLower(typename)) {
			if idx == index {
				return vw, true
			}
			idx += 1
		}
	}
	return View{}, false
}
