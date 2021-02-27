package ios

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

// List converts the parsed devices to the common format
func (d Device) List() types.Pairs {
	p := types.Pairs{}
	for _, devices := range d.Config.Devices {
		for _, device := range devices {
			if device.AvailabilityError != "" {
				log.Debugf("Device unavailable %s", device.Name)

				continue
			}
			var s = types.StateShutdown
			if device.State == "Booted" {
				s = types.StateBooted
			}
			p[device.Name] = types.Instance{
				ID:    device.UDID,
				State: s,
				Kind:  types.KindIos,
			}
		}
	}

	return p
}
