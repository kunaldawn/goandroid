// TODO : Documentation

package view

import (
	"errors"
	"fmt"
	"time"
)

func (devView DeviceView) IsTextPresent(text string, index int, timeout int) error {
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
		_, found := vws.GetByText(text, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after [%d] seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) IsMatchingTextPresnt(text string, index int, timeout int) error {
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
		_, found := vws.GetByMatchingText(text, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) ClickText(text string, index int, timeout int) error {
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
		vw, found := vws.GetByText(text, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) ClickMatchingText(text string, index int, timeout int) error {
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
		vw, found := vws.GetByMatchingText(text, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) GetViewForText(text string, index int, timeout int) (View, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return View{}, err
		}
		vw, found := vws.GetByText(text, index)
		if found {
			return vw, nil
		}
	}
	return View{}, errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) GetViewForMatchingText(text string, index int, timeout int) (View, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return View{}, err
		}
		vw, found := vws.GetByMatchingText(text, index)
		if found {
			return vw, nil
		}
	}
	return View{}, errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) ScrollDownToText(text string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsTextPresent(text, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Text [%s] not found after scrolling down [%d] times ", text, maxscroll))
}

func (devView DeviceView) ScrollUpToText(text string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsTextPresent(text, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Text [%s] not found after scrolling up [%d] times ", text, maxscroll))
}

func (devView DeviceView) ScrollDownToMatchingText(text string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTextPresnt(text, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Matching text [%s] not found after scrolling down [%d] times ", text, maxscroll))
}

func (devView DeviceView) ScrollUpToMatchingText(text string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTextPresnt(text, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Matching text [%s] not found after scrolling up [%d] times ", text, maxscroll))
}

func (devView DeviceView) GetTextForResource(resource string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByResource(resource, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) GetTextForMatchingResource(resource string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByMatchingResource(resource, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matcnhing resource [%s]", timeout, resource))
}

func (devView DeviceView) GetTextForType(typename string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByType(typename, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for type [%s]", timeout, typename))
}

func (devView DeviceView) GetTextForMatchingType(typename string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByMatchingType(typename, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching type [%s]", timeout, typename))
}

func (devView DeviceView) GetTextForDescription(description string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByDescription(description, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) GetTextForMatchingDescription(description string, index int, timeout int) (string, error) {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return "", err
		}
		vw, found := vws.GetByMatchingDescription(description, index)
		if found {
			return vw.Text, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}
