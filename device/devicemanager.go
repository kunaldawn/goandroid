// TODO : Documentation

package device

import (
	"github.com/kunaldawn/go-android/adbutility"
	"sync"
	"time"
)

type DeviceManager struct {
	Timeout int             // Timeout for all adb and shell operations
	ticker  *time.Ticker    // Ticker to update list
	devices map[string]bool // list of devices detected and its live status
	lock    sync.Mutex
}

func NewDeviceManager(refreshInterval int, timeout int) DeviceManager {
	ticker := time.NewTicker(time.Second * time.Duration(refreshInterval))
	return DeviceManager{devices: make(map[string]bool), Timeout: timeout, ticker: ticker, lock: sync.Mutex{}}
}

func (manager DeviceManager) StartUpdater() {
	go func() {
		err := manager.Refresh()
		if err != nil {
			panic(err)
		}
		for _ = range manager.ticker.C {
			err := manager.Refresh()
			if err != nil {
				panic(err)
			}
		}
	}()
}

func (manager DeviceManager) StopUpdater() {
	manager.ticker.Stop()
}

func (manager DeviceManager) Refresh() error {
	devices, err := adbutility.GetAttachedDevices(manager.Timeout)
	if err != nil {
		return err
	}
	manager.lock.Lock()
	for dev := range manager.devices {
		manager.devices[dev] = false
	}
	for index := range devices {
		manager.devices[devices[index]] = true
	}
	manager.lock.Unlock()
	return nil
}

func (manager DeviceManager) GetAllAvailableDevices() []Device {
	devs := []Device{}
	manager.lock.Lock()
	for dev := range manager.devices {
		if manager.devices[dev] {
			devs = append(devs, Device{Serial: dev, Timeout: manager.Timeout})
		}
	}
	manager.lock.Unlock()
	return devs
}

func (manager DeviceManager) GetAllDiscoveredDevices() []Device {
	devs := []Device{}
	manager.lock.Lock()
	for dev := range manager.devices {
		devs = append(devs, Device{Serial: dev, Timeout: manager.Timeout})
	}
	manager.lock.Unlock()
	return devs
}
