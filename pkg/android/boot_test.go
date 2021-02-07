package android

import (
	"testing"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/golang/mock/gomock"
)

func TestDevice_Boot(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Start().Return(nil)
	mockExecutable.EXPECT().String().Return("")

	if err := d.Boot(types.Instance{
		ID:    "Pixel_API_30",
		DebugPort: "5555",
		State: types.StateShutdown,
		Kind:  types.KindAndroid,
	}); err != nil {
		t.Error(err)
	}
}
