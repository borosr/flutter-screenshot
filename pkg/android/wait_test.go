package android

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/golang/mock/gomock"
)

func TestDevice_WaitUntilBooted(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	var wout *bytes.Buffer
	gomock.InOrder(
		mockExecutable.EXPECT().Run().Do(func() error {
			wout.WriteString("offline")
			return nil
		}),
		mockExecutable.EXPECT().Run().Do(func() error {
			wout.Reset()
			wout.WriteString(deviceState)
			return nil
		}),
		)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		wout = w.(*bytes.Buffer)
	})

	if err := d.WaitUntilBooted(types.Instance{}); err != nil {
		t.Error(err)
	}
}

func TestDevice_WaitUntilFailRunOnce(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	gomock.InOrder(
		mockExecutable.EXPECT().Run().Return(errors.New("failed")),
		mockExecutable.EXPECT().Run().Return(nil),
	)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte(deviceState))
	})

	if err := d.WaitUntilBooted(types.Instance{}); err != nil {
		t.Error(err)
	}
}

func TestDevice_WaitUntilFailMaxTimes(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().MaxTimes(maxErrorCount + 1).Return(errors.New("failed"))
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte("offline"))
	})

	if err := d.WaitUntilBooted(types.Instance{}); err == nil {
		t.Error("Missing error")
	} else if !errors.Is(err, ErrWaitingTooManyErrors) {
		t.Errorf("Error should be %v, instead of %v", ErrWaitingTooManyErrors, err)
	}
}
