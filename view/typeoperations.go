package view

import (
	"errors"
	"fmt"
	"time"
)

func (devView DeviceView) IsTypePresent(typename string, index int, timeout int) error {
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
		_, found := vws.GetByType(typename, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after [%d] seconds while searching for type [%s]", timeout, typename))
}

func (devView DeviceView) IsMatchingTypePresnt(typename string, index int, timeout int) error {
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
		_, found := vws.GetByMatchingType(typename, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching type [%s]", timeout, typename))
}

func (devView DeviceView) ClickType(typename string, index int, timeout int) error {
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
		vw, found := vws.GetByType(typename, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for type [%s]", timeout, typename))
}

func (devView DeviceView) ClickMatchingType(typename string, index int, timeout int) error {
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
		vw, found := vws.GetByMatchingType(typename, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching type [%s]", timeout, typename))
}

func (devView DeviceView) ScrollDownToType(typename string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsTypePresent(typename, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Type [$s] not found after scrolling down [%d] times ", typename, maxscroll))
}

func (devView DeviceView) ScrollUpToType(typename string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsTypePresent(typename, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Type [$s] not found after scrolling up [%d] times ", typename, maxscroll))
}

func (devView DeviceView) ScrollDownToMatchingType(typename string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTypePresnt(typename, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Matching type [$s] not found after scrolling down [%d] times ", typename, maxscroll))
}

func (devView DeviceView) ScrollUpToMatchingType(typename string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTypePresnt(typename, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Matching type [$s] not found after scrolling up [%d] times ", typename, maxscroll))
}

func (devView DeviceView) GetTypeForText(text string, index int, timeout int) (string, error) {
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
		vw, found := vws.GetByText(text, index)
		if found {
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) GetTypeForMatchingText(text string, index int, timeout int) (string, error) {
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
		vw, found := vws.GetByMatchingText(text, index)
		if found {
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matcnhing text [%s]", timeout, text))
}

func (devView DeviceView) GetTypeForResource(resource string, index int, timeout int) (string, error) {
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
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) GetTypeForMatchingResource(resource string, index int, timeout int) (string, error) {
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
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching resource [%s]", timeout, resource))
}

func (devView DeviceView) GetTypeForDescription(description string, index int, timeout int) (string, error) {
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
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) GetTypeForMatchingDescription(description string, index int, timeout int) (string, error) {
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
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}
