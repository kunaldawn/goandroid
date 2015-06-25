// TODO : Documentation

package view

import (
	"errors"
	"fmt"
	"github.com/kunaldawn/goandroid/input"
	"github.com/kunaldawn/goandroid/logging"
	"time"
)

func (devView DeviceView) IsTextPresent(text string, index int, timeout int) error {
	logging.Log("IsTextPresent : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("IsMatchingTextPresnt : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
	logging.Log("ClickText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for text [%s]", timeout, text))
}

func (devView DeviceView) ClickMatchingText(text string, index int, timeout int) error {
	logging.Log("ClickMatchingText : text [%s] : index [%d] : timeout [%d]", text, index, timeout)
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
			im := input.NewInputManager(devView.dev)
			return im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for matching text [%s]", timeout, text))
}

func (devView DeviceView) GetTextForResource(resource string, index int, timeout int) (string, error) {
	logging.Log("GetTextForResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("GetTextForMatchingResource : resource [%s] : index [%d] : timeout [%d]", resource, index, timeout)
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
	logging.Log("GetTextForType : typename [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
	logging.Log("GetTextForMatchingType : typename [%s] : index [%d] : timeout [%d]", typename, index, timeout)
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
	logging.Log("GetTextForDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
	logging.Log("GetTextForMatchingDescription : description [%s] : index [%d] : timeout [%d]", description, index, timeout)
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
