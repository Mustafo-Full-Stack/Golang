package models

import (
	"errors"
	"strings"
	"testing"
)

func TestRegisterInput(t *testing.T) {
	valid := RegisterInput{Email: "ivan@mail.com", Password: "secret123", Name: "Иван"}

	tests := []struct {
		name    string
		change  func(in *RegisterInput)
		wantErr bool
	}{
		{name: "valid", change: func(*RegisterInput) {}},
		{name: "email is normalized", change: func(in *RegisterInput) { in.Email = "  IVAN@Mail.com " }},
		{name: "invalid email", change: func(in *RegisterInput) { in.Email = "ivan" }, wantErr: true},
		{name: "email with display name", change: func(in *RegisterInput) { in.Email = "Ivan <ivan@mail.com>" }, wantErr: true},
		{name: "empty name", change: func(in *RegisterInput) { in.Name = "   " }, wantErr: true},
		{name: "long name", change: func(in *RegisterInput) { in.Name = strings.Repeat("a", maxNameRunes+1) }, wantErr: true},
		{name: "nul in name", change: func(in *RegisterInput) { in.Name = "a\x00b" }, wantErr: true},
		{name: "short password", change: func(in *RegisterInput) { in.Password = "1234567" }, wantErr: true},
		{name: "long password", change: func(in *RegisterInput) { in.Password = strings.Repeat("a", maxPasswordBytes+1) }, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := valid
			tt.change(&in)
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
			if in.Email != "ivan@mail.com" {
				t.Fatalf("Email = %q, want normalized", in.Email)
			}
		})
	}
}
