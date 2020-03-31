package core

import (
	"testing"
)

func TestMultipliedByE8_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		want    MultipliedBy1e8
		args    string
		wantErr bool
	}{{
		args:    "6610.29000000",
		want:    661029000000,
		wantErr: false,
	}, {
		args:    "\"6610.29\"",
		want:    661029000000,
		wantErr: false,
	}, {
		args:    ".29",
		want:    29000000,
		wantErr: false,
	}, {
		args:    "1",
		want:    100000000,
		wantErr: false,
	}, {
		args:    "\"0.2.9\"",
		wantErr: true,
	}, {
		args:    "",
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			result := new(MultipliedBy1e8)
			err := result.UnmarshalJSON([]byte(tt.args))
			if (err != nil) != tt.wantErr {
				t.Fatalf("MultipliedByE8.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if *result != tt.want {
				t.Errorf("MultipliedByE8.UnmarshalJSON() result = %v, want %v", *result, tt.want)
			}
		})
	}
}
