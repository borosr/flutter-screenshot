package executor

import (
	"errors"
	"reflect"
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/exec"
	"github.com/borosr/flutter-screenshot/src/config"
	. "github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func TestExecuteExistingDevice(t *testing.T) {
	for _, theme := range config.AllModes {
		ctrl := gomock.NewController(t)
		mockDeviceAction := NewMockDeviceAction(ctrl)
		deviceID := gofakeit.UUID()
		instance := Instance{
			ID:    deviceID,
			State: StateShutdown,
			Kind:  KindIos,
		}
		mockDeviceAction.EXPECT().List().Return(Pairs{
			"iPhone X": instance,
		})
		mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
		mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(nil)
		gomock.InOrder(
			mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Any()).Return(nil),
			mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Any()).Return(nil),
		)
		mockDeviceAction.EXPECT().Shutdown(gomock.Eq(instance)).Return(nil)

		if err := execute([]config.Device{{
			Name: "iPhone X",
			Mode: theme,
		}}, `echo "hello"`, mockDeviceAction); err != nil {
			t.Error(err)
		}
	}
}

func TestExecuteNotExistingDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	instance := Instance{
		ID:    deviceID,
		State: StateBooted,
		Kind:  KindIos,
	}
	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
	mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Eq(config.ModeLight.String())).Return(nil)
	mockDeviceAction.EXPECT().Shutdown(gomock.Eq(instance)).Return(nil)

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, `echo "hello"`, mockDeviceAction); err != nil {
		t.Error(err)
	}
}

func TestExecuteDeviceCreationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceCreateError := errors.New("unable to create device")

	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return("", KindIos, deviceCreateError)

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, "", mockDeviceAction); err == nil {
		t.Error("missing error")
	} else if !errors.Is(err, deviceCreateError) {
		t.Errorf("error should be %v, instead of %v", deviceCreateError, err)
	}
}

func TestExecuteDeviceBootError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	instance := Instance{
		ID:    deviceID,
		State: StateBooted,
		Kind:  KindIos,
	}
	bootError := errors.New("unable to boot device")

	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
	mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(bootError)

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, "", mockDeviceAction); err == nil {
		t.Error("missing error")
	} else if !errors.Is(err, bootError) {
		t.Errorf("error should be %v, instead of %v", bootError, err)
	}
}

func TestExecuteDeviceWaitForBootedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	instance := Instance{
		ID:    deviceID,
		State: StateBooted,
		Kind:  KindIos,
	}
	waitForBootedError := errors.New("unable wait for device")

	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
	mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(waitForBootedError)

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, "", mockDeviceAction); err == nil {
		t.Error("missing error")
	} else if !errors.Is(err, waitForBootedError) {
		t.Errorf("error should be %v, instead of %v", waitForBootedError, err)
	}
}

func TestExecuteDeviceSetThemeError(t *testing.T) {
	for _, theme := range append(config.AllModes, "") {
		ctrl := gomock.NewController(t)
		mockDeviceAction := NewMockDeviceAction(ctrl)

		const deviceName = "iPhone X"
		deviceID := gofakeit.UUID()
		instance := Instance{
			ID:    deviceID,
			State: StateBooted,
			Kind:  KindIos,
		}
		setThemeError := errors.New("unable set device's theme")

		mockDeviceAction.EXPECT().List().Return(Pairs{})
		mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
		mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
		mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(nil)
		gomock.InOrder(
			mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Any()).Return(setThemeError),
			mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Any()).Return(setThemeError),
		)
		mockDeviceAction.EXPECT().Shutdown(gomock.Eq(instance)).Return(nil)

		if err := execute([]config.Device{{
			Name: deviceName,
			Mode: theme,
		}}, "", mockDeviceAction); err == nil {
			t.Error("missing error")
		} else if !errors.Is(err, setThemeError) {
			t.Errorf("error should be %v, instead of %v", setThemeError, err)
		}
	}
}

func TestExecuteDeviceShutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	instance := Instance{
		ID:    deviceID,
		State: StateBooted,
		Kind:  KindIos,
	}
	shutdownError := errors.New("unable shutdown device")

	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
	mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Eq(config.ModeLight.String())).Return(nil)
	mockDeviceAction.EXPECT().Shutdown(gomock.Eq(instance)).Return(shutdownError)

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, "", mockDeviceAction); err == nil {
		t.Error("missing error")
	} else if !errors.Is(err, shutdownError) {
		t.Errorf("error should be %v, instead of %v", shutdownError, err)
	}
}

func TestExecuteDeviceCommandExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDeviceAction := NewMockDeviceAction(ctrl)

	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	instance := Instance{
		ID:    deviceID,
		State: StateBooted,
		Kind:  KindIos,
	}
	cmdExecuteError := errors.New("command execution error")

	mockExecutable := NewMockExecutable(ctrl)
	invoke = mockExecute(mockExecutable)

	mockDeviceAction.EXPECT().List().Return(Pairs{})
	mockDeviceAction.EXPECT().Create(gomock.Eq(deviceName)).Return(deviceID, KindIos, nil)
	mockDeviceAction.EXPECT().Boot(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().WaitUntilBooted(gomock.Eq(instance)).Return(nil)
	mockDeviceAction.EXPECT().SetTheme(gomock.Eq(instance), gomock.Eq(config.ModeLight.String())).Return(nil)
	mockDeviceAction.EXPECT().Shutdown(gomock.Eq(instance)).Return(nil)
	mockExecutable.EXPECT().Run().Return(cmdExecuteError)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).Return()

	if err := execute([]config.Device{{
		Name: deviceName,
		Mode: config.ModeLight.String(),
	}}, "", mockDeviceAction); err == nil {
		t.Error("missing error")
	} else if !errors.Is(err, cmdExecuteError) {
		t.Errorf("error should be %v, instead of %v", cmdExecuteError, err)
	}
}

func TestSplitCommand(t *testing.T) {
	type output struct {
		cmd     string
		subCmds []string
	}
	cases := []struct {
		name string
		in   string
		out  output
	}{
		{
			name: "empty string",
			in:   "",
			out: output{
				cmd:     "",
				subCmds: []string{},
			},
		},
		{
			name: "single command",
			in:   "echo",
			out: output{
				cmd:     "echo",
				subCmds: []string{},
			},
		},
		{
			name: "single command multiple subcommands",
			in:   "echo asd asd asd",
			out: output{
				cmd:     "echo",
				subCmds: []string{"asd", "asd", "asd"},
			},
		},
		{
			name: "single command multiple subcommands and flags",
			in:   "echo asd asd --help -n name",
			out: output{
				cmd:     "echo",
				subCmds: []string{"asd", "asd", "--help", "-n", "name"},
			},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			cmd, subCmds := splitCommand(c.in)
			if cmd != c.out.cmd {
				t.Errorf("The cmd output should be %s, instead of %s", c.out.cmd, cmd)
			}
			if !reflect.DeepEqual(subCmds, c.out.subCmds) {
				t.Errorf("The sub commands output should be %v, instead of %v", c.out.subCmds, subCmds)
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	invoke = exec.Command
	if err := executeCommand("echo test 123", "some_uuid"); err != nil {
		t.Error(err)
	}
}
