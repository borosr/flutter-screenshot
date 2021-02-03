package ios

import (
	"testing"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/golang/mock/gomock"
)

func TestDevice_BootShutDowned(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")

	if err := d.Boot(types.Instance{
		ID:    "",
		State: types.StateShutdown,
		Kind:  0,
	}); err != nil {
		t.Error(err)
	}
}

func TestDevice_BootBooted(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	gomock.InOrder(
		mockExecutable.EXPECT().Run().Return(nil),
		mockExecutable.EXPECT().Run().Return(nil),
	)
	gomock.InOrder(
		mockExecutable.EXPECT().String().Return(""),
		mockExecutable.EXPECT().String().Return(""),
	)

	if err := d.Boot(types.Instance{
		ID:    "",
		State: types.StateBooted,
		Kind:  0,
	}); err != nil {
		t.Error(err)
	}
}
