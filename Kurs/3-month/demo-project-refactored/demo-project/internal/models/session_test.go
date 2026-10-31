package models

import (
	"errors"
	"testing"
)

func TestLoginInputDeviceType(t *testing.T) {
	tests := []struct {
		device  DeviceType
		want    DeviceType
		wantErr bool
	}{
		{device: "", want: DeviceWeb},
		{device: "web", want: DeviceWeb},
		{device: " Mobile ", want: DeviceMobile},
		{device: "tv", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(string(tt.device), func(t *testing.T) {
			in := LoginInput{Email: "ivan@mail.com", Password: "secret123", DeviceType: tt.device}
			err := in.Validate()
			if tt.wantErr {
				if !errors.Is(err, ErrInvalidInput) {
					t.Fatalf("Validate() error = %v, want ErrInvalidInput", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Validate() error = %v", err)
			}
			if in.DeviceType != tt.want {
				t.Fatalf("DeviceType = %q, want %q", in.DeviceType, tt.want)
			}
		})
	}
}

func TestLoginInputRequiresEmailAndPassword(t *testing.T) {
	in := LoginInput{Email: "  ", Password: "secret123"}
	if err := in.Validate(); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("Validate() error = %v, want ErrInvalidInput", err)
	}
}
