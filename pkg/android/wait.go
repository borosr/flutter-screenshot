package android

import (
	"bytes"
	"errors"
	"time"

	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

const (
	maxErrorCount = 10
	deviceState   = "device"
)

var (
	ErrWaitingTooManyErrors = errors.New("too many errors got when waiting to startup")
)

func (d Device) WaitUntilBooted(instance types.Instance) error {
	cmd := execute("adb", "-s", "emulator-"+instance.DebugPort, "get-state")
	log.Debugf("Emulator state: Executing cmd: %s", cmd.String())
	var out bytes.Buffer
	cmd.Stdout(&out)
	var ec int
	var start = time.Now()
	for {
		if ec >= maxErrorCount {
			return ErrWaitingTooManyErrors
		}
		if err := cmd.Run(); err != nil {
			log.Error(err)
			ec++
			time.Sleep(time.Second)

			continue
		}

		if out.String() == deviceState {
			return nil
		}
		time.Sleep(time.Second)
		log.Debugf("Waiting since %v", time.Since(start))
	}
}
