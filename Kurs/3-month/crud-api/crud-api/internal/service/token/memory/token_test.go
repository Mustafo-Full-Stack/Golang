package memory

import (
	"context"
	"crud-api/internal/models"
	"errors"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	ctx := context.Background()
	m := NewTokenManager(time.Hour)

	value, err := m.Generate(ctx, 42)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}

	userID, err := m.Validate(ctx, value)
	if err != nil {
		t.Fatalf("validate: %v", err)
	}
	if userID != 42 {
		t.Fatalf("userID = %d, want 42", userID)
	}
}

func TestValidateExpired(t *testing.T) {
	ctx := context.Background()
	m := NewTokenManager(time.Millisecond)

	value, err := m.Generate(ctx, 42)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	time.Sleep(5 * time.Millisecond)

	if _, err := m.Validate(ctx, value); !errors.Is(err, models.ErrInvalidToken) {
		t.Fatalf("err = %v, want %v", err, models.ErrInvalidToken)
	}
}

func TestRevoke(t *testing.T) {
	ctx := context.Background()
	m := NewTokenManager(time.Hour)

	value, err := m.Generate(ctx, 42)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if err := m.Revoke(ctx, value); err != nil {
		t.Fatalf("revoke: %v", err)
	}

	if _, err := m.Validate(ctx, value); !errors.Is(err, models.ErrInvalidToken) {
		t.Fatalf("err = %v, want %v", err, models.ErrInvalidToken)
	}
}
