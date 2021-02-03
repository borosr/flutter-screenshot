package ios

import (
	"testing"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/golang/mock/gomock"
)

func TestDevice_SetTheme(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).Return()

	if err := d.SetTheme(types.Instance{}, "light"); err != nil {
		t.Error(err)
	}
}
