// Package view provides various query and UI operation methods on android
// UI View. It internally uses uiautomator XML dump mechanism to parse device
// view hierarchy, so all operations are based on information extracted from
// XML hierarchy dump.
// TODO : Detailed package documentation need to be done here.
package view

import (
	"github.com/kunaldawn/goandroid/geometry"
	"strings"
)

// View struct is an internal representation of android ui component. This is
// generated from uiautomaror XML dump but in a representation that is more
// suitable for various UI operations by goandroid framework. It is a typed
// representation of uiautomator node. Some extra fields are calculated for
// ease of use, such as center of the view, which allows to click the view
// and index of the view. Index is calculated based on occurence of the
// node in uiautomator xml dump.
type View struct {
	Index         int            // Index value of the android view component.
	Class         string         // Canonical class name of the android view component.
	Package       string         // Canonical package name of the android view component.
	Resource      string         // Associated resource id (if any) of the view component.
	Text          string         // Associated text (if any) of the view component.
	Description   string         // Associated description (if any) of the view component.
	Clickable     bool           // Boolean value indicating if the view is clickable.
	Checkable     bool           // Boolean value indicating if the view is checkbox or not.
	Checked       bool           // Boolean value indicating if the view is checked or not.
	Enabled       bool           // Boolean value indicating if the view is enabled or not.
	Focusable     bool           // Boolean value indicating if the view is focusable by user or not.
	Focused       bool           // Boolean value indicating if the view is currently focused or not.
	Scrollable    bool           // Boolean value indicating if the view is scrollable or not.
	LongClickable bool           // Boolean value indicating if the view is long click enabled or not.
	Password      bool           // Boolean value indicating if the view is a password field or not.
	Selected      bool           // Boolean value indicating if the view is currently selected by user or not.
	Bound         geometry.Rect  // Bounding rectangle of the view.
	Center        geometry.Point // Center coordinate of the view.
}

// Views is a type that represents sclice of View structure.
type Views []View

// GetByText method returns a View struct based on excat match of specified
// text and index. It also returns a boolean value indicating if the match
// is found or not. Please note that parameter index is zero based index
// system, that is, index value of first element is zero. If no view is
// found with exact text match, empty view is returned with boolean value
// false indicating view is not found.
func (views Views) GetByText(text string, index int) (View, bool) {
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

// GetByMatchingText method returns a View struct on matching text with
// specified index. It also returns a boolean value indicating if the match
// is found or not. Please note that parameter index is zero based index
// system, that is, index value of first element is zero. If no view is
// found with matching text, empty view is returned with boolean value
// false indicating view is not found. Text match is case insensitive.
func (views Views) GetByMatchingText(text string, index int) (View, bool) {
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

// GetByResource method returns a View struct based on excat match of
// specified text value of view resource id and index. It also returns a
// boolean value indicating if the match is found or not. Please note that
// parameter index is zero based index system, that is, index value of first
// element is zero. If no view is found with exact resource id match, empty
// view is returned with boolean value false indicating view is not found.
// NOTE : Here resource id represents following "<resource_id>" part only
//        <package_name>:id/<resource_id>
func (views Views) GetByResource(resource string, index int) (View, bool) {
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

// GetByMatchingResource method returns a View struct on matching text value
// of resource id with specified index. It also returns a boolean value
// indicating if the match is found or not. Please note that parameter index
// is zero based index system, that is, index value of first element is zero.
// If no view is found with matching test value of resource id, empty view is
// returned with boolean value false indicating view is not found. Resource
// id match is case insensitive.
// NOTE : Here resource id represents following "<resource_id>" part only
//        <package_name>:id/<resource_id>
func (views Views) GetByMatchingResource(resource string, index int) (View, bool) {
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

// GetByDescription method returns a View struct based on excat match of
// specified text value of views description and index. It also returns a
// boolean value indicating if the match is found or not. Please note that
// parameter index is zero based index system, that is, index value of first
// element is zero. If no view is found with exact description text match,
// empty view is returned with boolean value false indicating view is not found.
func (views Views) GetByDescription(description string, index int) (View, bool) {
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

// GetByMatchingDescription method returns a View struct on matching text
// value of description field of the view with specified index. It also
// returns a boolean value indicating if the match is found or not. Please
// note that parameter index is zero based index system, that is, index value
// of first element is zero. If no view is found with matching description
// value, empty view is returned with boolean value false indicating view is
// not found. Description match is case insensitive.
func (views Views) GetByMatchingDescription(description string, index int) (View, bool) {
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

// GetByType method returns a View struct based on excat match of specified
// text value of views class name and index. It also returns a boolean value
// indicating if the match is found or not. Please note that parameter index
// is zero based index system, that is, index value of first element is zero.
// If no view is found with exact class name match, empty view is returned
// with boolean value false indicating view is not found.
func (views Views) GetByType(typename string, index int) (View, bool) {
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

// GetByMatchingType method returns a View struct on matching text value of
// class name field of the view with specified index. It also returns a
// boolean value indicating if the match is found or not. Please note that
// parameter index is zero based index system, that is, index value of first
// element is zero. If no view is found with matching class name, empty view
// is returned with boolean value false indicating view is not found. Class
// name match is case insensitive.
func (views Views) GetByMatchingType(typename string, index int) (View, bool) {
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
