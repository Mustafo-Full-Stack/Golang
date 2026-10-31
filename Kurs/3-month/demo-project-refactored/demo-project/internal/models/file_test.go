package models

import (
	"errors"
	"strings"
	"testing"
)

func TestUploadFileInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "keeps name", input: "photo.png", want: "photo.png"},
		{name: "strips path", input: "../../etc/passwd", want: "passwd"},
		{name: "trims", input: "  photo.png  ", want: "photo.png"},
		{name: "empty", input: "   ", wantErr: true},
		{name: "too long", input: strings.Repeat("a", maxFileNameRunes+1), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := UploadFileInput{Name: tt.input, Content: strings.NewReader("x")}
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
			if in.Name != tt.want {
				t.Fatalf("Name = %q, want %q", in.Name, tt.want)
			}
		})
	}
}

func TestUploadFileInputRequiresContent(t *testing.T) {
	in := UploadFileInput{Name: "photo.png"}
	if err := in.Validate(); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("Validate() error = %v, want ErrInvalidInput", err)
	}
}
