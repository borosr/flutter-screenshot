package ios

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestDevice_Shutdown(t *testing.T) {
	d := Device{}
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")

	if err := d.Shutdown(""); err != nil {
		t.Error(err)
	}
}
