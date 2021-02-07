package types

import "testing"

func TestPairs_String(t *testing.T) {
	tests := []struct {
		name string
		p    Pairs
		want string
	}{
		{
			name: "null value",
			p:    nil,
			want: "null",
		},
		{
			name: "basic parse test",
			p:    Pairs{
				"key": Instance{
					ID:    "id",
					State: 1,
					Kind:  1,
				},
			},
			want: `{"key":{"id":"id","debug_port":"","state":1,"kind":1}}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
