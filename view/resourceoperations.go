package view

import (
	"errors"
	"fmt"
	"time"
)

func (devView DeviceView) IsResourcePresent(resource string, index int, timeout int) error {
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
		_, found := vws.GetByResource(resource, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after [%d] seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) IsMatchingResourcePresnt(resource string, index int, timeout int) error {
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
		_, found := vws.GetByMatchingResource(resource, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching resource [%s]", timeout, resource))
}

func (devView DeviceView) ClickResource(resource string, index int, timeout int) error {
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
		vw, found := vws.GetByResource(resource, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) ClickMatchingResource(resource string, index int, timeout int) error {
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
		vw, found := vws.GetByMatchingResource(resource, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching resource [%s]", timeout, resource))
}

func (devView DeviceView) ScrollDownToResource(resource string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsResourcePresent(resource, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Resource [$s] not found after scrolling down [%d] times ", resource, maxscroll))
}

func (devView DeviceView) ScrollUpToResource(resource string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsResourcePresent(resource, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Resource [$s] not found after scrolling up [%d] times ", resource, maxscroll))
}

func (devView DeviceView) ScrollDownToMatchingResource(resource string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTextPresnt(resource, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Matching resource [$s] not found after scrolling down [%d] times ", resource, maxscroll))
}

func (devView DeviceView) ScrollUpToMatchingResource(resource string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingTextPresnt(resource, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Matching resource [$s] not found after scrolling up [%d] times ", resource, maxscroll))
}

func (devView DeviceView) GetResourceForText(text string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) GetResourceForMatchingText(text string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matcnhing text [%s]", timeout, text))
}

func (devView DeviceView) GetResourceForType(typename string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for type [%s]", timeout, typename))
}

func (devView DeviceView) GetResourceForMatchingType(typename string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching type [%s]", timeout, typename))
}

func (devView DeviceView) GetResourceForDescription(description string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) GetResourceForMatchingDescription(description string, index int, timeout int) (string, error) {
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
			return vw.Resource, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}
