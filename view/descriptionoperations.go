package view

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/logging"
	"time"
)

func (devView DeviceView) IsDescriptionPresent(description string, index int, timeout int) error {
	logging.Log("IsDescriptionPresent : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
	logging.Log("IsMatchingDescriptionPresnt : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
	logging.Log("ClickDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for description [%s]", timeout, description))
}

func (devView DeviceView) ClickMatchingDescription(description string, index int, timeout int) error {
	logging.Log("ClickMatchingDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching description [%s]", timeout, description))
}

func (devView DeviceView) GetDescriptionForText(text string, index int, timeout int) (string, error) {
	logging.Log("GetTypeForText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("GetTypeForMatchingText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("GetTypeForResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("GettypeForMatchingResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("GetTypeForDescription : type [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
	logging.Log("GetTypeForMatchingDescription : type [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
