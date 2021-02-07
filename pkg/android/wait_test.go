package android

import (
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

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte(deviceState))
	})

	if err := d.WaitUntilBooted(types.Instance{}); err != nil {
		t.Error(err)
	}
}
