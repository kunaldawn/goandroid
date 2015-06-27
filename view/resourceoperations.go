package view

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/logging"
	"time"
)

func (devView DeviceView) IsResourcePresent(resource string, index int, timeout int) error {
	logging.Log("IsResourcePresent : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("IsMatchingResourcePresnt : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("ClickResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

func (devView DeviceView) ClickMatchingResource(resource string, index int, timeout int) error {
	logging.Log("ClickMatchingResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching resource [%s]", timeout, resource))
}

func (devView DeviceView) GetResourceForText(text string, index int, timeout int) (string, error) {
	logging.Log("GetResourceForText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("GetResourceForMatchingText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("GetResourceForType : type [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
	logging.Log("GetResourceForMatchingType : type [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
	logging.Log("GetResourceForDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
	logging.Log("GetResourceForMatchingDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
