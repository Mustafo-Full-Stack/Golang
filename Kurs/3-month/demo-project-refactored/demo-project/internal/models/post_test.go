package models

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestCreatePostInput(t *testing.T) {
	tests := []struct {
		name    string
		in      CreatePostInput
		wantErr bool
	}{
		{name: "valid", in: CreatePostInput{Title: "Заголовок", Content: "Текст"}},
		{name: "trims spaces", in: CreatePostInput{Title: "  Заголовок  ", Content: "  Текст  "}},
		{name: "empty title", in: CreatePostInput{Title: " ", Content: "Текст"}, wantErr: true},
		{name: "empty content", in: CreatePostInput{Title: "Заголовок", Content: " "}, wantErr: true},
		{name: "nul in title", in: CreatePostInput{Title: "a\x00b", Content: "Текст"}, wantErr: true},
		{name: "nul in content", in: CreatePostInput{Title: "Заголовок", Content: "a\x00b"}, wantErr: true},
		{name: "title too long", in: CreatePostInput{Title: strings.Repeat("a", maxTitleRunes+1), Content: "Текст"}, wantErr: true},
		{name: "content too long", in: CreatePostInput{Title: "Заголовок", Content: strings.Repeat("a", maxContentRunes+1)}, wantErr: true},
		{name: "negative file id", in: CreatePostInput{Title: "a", Content: "b", FileIDs: []int64{1, -2}}, wantErr: true},
		{name: "too many files", in: CreatePostInput{Title: "a", Content: "b", FileIDs: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.in
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
		})
	}
}

func TestCreatePostInputDeduplicatesFiles(t *testing.T) {
	in := CreatePostInput{Title: "Заголовок", Content: "Текст", FileIDs: []int64{1, 2, 1, 3}}
	if err := in.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if want := []int64{1, 2, 3}; !slices.Equal(in.FileIDs, want) {
		t.Fatalf("FileIDs = %v, want %v", in.FileIDs, want)
	}
}

func TestUpdatePostInputRequiresID(t *testing.T) {
	in := UpdatePostInput{Title: "Заголовок", Content: "Текст"}
	if err := in.Validate(); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("Validate() error = %v, want ErrInvalidInput", err)
	}
}
