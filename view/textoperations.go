// TODO : Documentation

package view

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/input"
	"time"
)

func (devView DeviceView) IsTextPresent(text string, timeout int) error {
	return devView.IsTextPresentIndex(text, 0, timeout)
}

func (devView DeviceView) IsMatchingTextPresnt(text string, timeout int) error {
	return devView.IsMatchingTextPresntIndex(text, 0, timeout)
}

func (devView DeviceView) ClickText(text string, timeout int) error {
	return devView.ClickTextIndex(text, 0, timeout)
}

func (devView DeviceView) ClickMatchingText(text string, timeout int) error {
	return devView.ClickMatchingTextIndex(text, 0, timeout)
}

func (devView DeviceView) GetTextForResource(resource string, timeout int) error {
	return devView.GetTextForResourceIndex(resource, 0, timeout)
}

func (devView DeviceView) GetTextForType(typename string, timeout int) error {
	return devView.GetTextForTypeIndex(typename, 0, timeout)
}

func (devView DeviceView) GetTextForDescription(description string, timeout int) error {
	return devView.GetTextForDescriptionIndex(description, 0, timeout)
}

func (devView DeviceView) IsTextPresentIndex(text string, index int, timeout int) error {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return err
		}
		_, found := vws.GetByTextIndex(text, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after [%d] seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) IsMatchingTextPresntIndex(text string, index int, timeout int) error {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return err
		}
		_, found := vws.GetByMatchingTextIndex(text, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) ClickTextIndex(text string, index int, timeout int) error {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return err
		}
		vw, found := vws.GetByTextIndex(text, index)
		if found {
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) ClickMatchingTextIndex(text string, index int, timeout int) error {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return err
		}
		vw, found := vws.GetByMatchingTextIndex(text, index)
		if found {
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) GetTextForResourceIndex(resource string, index int, timeout int) error {
	// TODO : Implement this method
	return nil
}

func (devView DeviceView) GetTextForTypeIndex(typename string, index int, timeout int) error {
	// TODO : Implement this method
	return nil
}

func (devView DeviceView) GetTextForDescriptionIndex(description string, index int, timeout int) error {
	// TODO : Implement this method
	return nil
}
