package core

import "testing"

func Test_sliceContainsNotNil(t *testing.T) {
	tests := []struct {
		name string
		args []func()
		want bool
	}{
		{
			name: "nil input",
			args: nil,
			want: false,
		},
		{
			name: "zero len",
			args: []func(){},
			want: false,
		},
		{
			name: "two nil",
			args: []func(){nil, nil},
			want: false,
		},
		{
			name: "two normal",
			args: []func(){func() {}, func() {}},
			want: true,
		},
		{
			name: "two normal and nil",
			args: []func(){func() {}, nil, func() {}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceContainsNotNil(tt.args); got != tt.want {
				t.Errorf("sliceContainsNotNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
