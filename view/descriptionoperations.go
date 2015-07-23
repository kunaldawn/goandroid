package view

import (
	"errors"
	"fmt"
	"time"
)

func (devView DeviceView) IsDescriptionPresent(description string, index int, timeout int) error {
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
		_, found := vws.GetByDescription(description, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after [%d] seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) IsMatchingDescriptionPresnt(description string, index int, timeout int) error {
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
		_, found := vws.GetByMatchingDescription(description, index)
		if found {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}

func (devView DeviceView) ClickDescription(description string, index int, timeout int) error {
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
		vw, found := vws.GetByDescription(description, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) ClickMatchingDescription(description string, index int, timeout int) error {
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
		vw, found := vws.GetByMatchingDescription(description, index)
		if found {
			return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}

func (devView DeviceView) GetViewForDescription(description string, index int, timeout int) (View, error) {
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
		vw, found := vws.GetByDescription(description, index)
		if found {
			return vw, nil
		}
	}
	return View{}, errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) GetViewForMatchingDescription(description string, index int, timeout int) (View, error) {
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
		vw, found := vws.GetByMatchingDescription(description, index)
		if found {
			return vw, nil
		}
	}
	return View{}, errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}

func (devView DeviceView) ScrollDownToDescription(description string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsDescriptionPresent(description, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Description [$s] not found after scrolling down [%d] times ", description, maxscroll))
}

func (devView DeviceView) ScrollUpToDescription(description string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsDescriptionPresent(description, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Description [$s] not found after scrolling up [%d] times ", description, maxscroll))
}

func (devView DeviceView) ScrollDownToMatchingDescription(description string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingDescriptionPresnt(description, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeUp(1)
	}
	return errors.New(fmt.Sprintf("Matching description [$s] not found after scrolling down [%d] times ", description, maxscroll))
}

func (devView DeviceView) ScrollUpToMatchingDescription(description string, index int, maxscroll int) error {
	for i := 0; i < maxscroll; i++ {
		err := devView.IsMatchingDescriptionPresnt(description, index, 1)
		if err == nil {
			return nil
		}
		devView.im.TouchScreen.SwipeDown(1)
	}
	return errors.New(fmt.Sprintf("Matching description [$s] not found after scrolling up [%d] times ", description, maxscroll))
}

func (devView DeviceView) GetDescriptionForText(text string, index int, timeout int) (string, error) {
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
			return vw.Description, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) GetDescriptionForMatchingText(text string, index int, timeout int) (string, error) {
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
			return vw.Description, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matcnhing text [%s]", timeout, text))
}

func (devView DeviceView) GetDescriptionForResource(resource string, index int, timeout int) (string, error) {
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
			return vw.Description, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) GetDescriptionForMatchingResource(resource string, index int, timeout int) (string, error) {
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
			return vw.Description, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching resource [%s]", timeout, resource))
}

func (devView DeviceView) GetDescriptionFortype(typename string, index int, timeout int) (string, error) {
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
			return vw.Description, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for type [%s]", timeout, typename))
}

func (devView DeviceView) GetDescriptionForMatchingType(typename string, index int, timeout int) (string, error) {
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
			return vw.Class, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching type [%s]", timeout, typename))
}
