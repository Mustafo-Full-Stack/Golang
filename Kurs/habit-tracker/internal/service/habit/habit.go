package habit

import (
	"context"
	"fmt"
	"time"

	"habit-tracker/internal/models"
)

const maxListLimit = 100

type Repository interface {
	Create(ctx context.Context, h *models.Habit) error
	ListByUser(ctx context.Context, userID int64, limit, offset int) ([]models.Habit, error)
	CountByUser(ctx context.Context, userID int64) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Habit, error)
	CreateCheck(ctx context.Context, habitID int64, date time.Time) error
	ListChecks(ctx context.Context, habitID int64, limit, offset int) ([]models.HabitCheck, error)
	CountChecks(ctx context.Context, habitID int64) (int64, error)
}

type Service struct {
	habits Repository
}

func NewService(habits Repository) *Service {
	return &Service{habits: habits}
}

func (s *Service) Create(ctx context.Context, in models.CreateHabitInput) (*models.Habit, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	habit := &models.Habit{UserID: in.UserID, Name: in.Name}

	if err := s.habits.Create(ctx, habit); err != nil {
		return nil, err
	}

	return habit, nil
}

func (s *Service) List(ctx context.Context, userID int64, limit, offset int) ([]models.Habit, int64, error) {
	if err := validatePagination(limit, offset); err != nil {
		return nil, 0, err
	}

	habits, err := s.habits.ListByUser(ctx, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.habits.CountByUser(ctx, userID)
	if err != nil {
		return nil, 0, err
	}

	return habits, total, nil
}

func (s *Service) Check(ctx context.Context, userID, habitID int64) error {
	if err := s.checkOwner(ctx, userID, habitID); err != nil {
		return err
	}

	now := time.Now().UTC()

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	return s.habits.CreateCheck(ctx, habitID, today)
}

func (s *Service) Checks(ctx context.Context, userID, habitID int64, limit, offset int) ([]models.HabitCheck, int64, error) {
	if err := validatePagination(limit, offset); err != nil {
		return nil, 0, err
	}

	if err := s.checkOwner(ctx, userID, habitID); err != nil {
		return nil, 0, err
	}

	checks, err := s.habits.ListChecks(ctx, habitID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.habits.CountChecks(ctx, habitID)
	if err != nil {
		return nil, 0, err
	}

	return checks, total, nil
}

func validatePagination(limit, offset int) error {
	if limit < 1 || limit > maxListLimit {
		return fmt.Errorf("%w: limit must be between 1 and %d", models.ErrInvalidInput, maxListLimit)
	}

	if offset < 0 {
		return fmt.Errorf("%w: offset must be zero or greater", models.ErrInvalidInput)
	}

	return nil
}

func (s *Service) checkOwner(ctx context.Context, userID, habitID int64) error {
	habit, err := s.habits.GetByID(ctx, habitID)

	if err != nil {
		return err
	}

	if habit.UserID != userID {
		return models.ErrForbidden
	}

	return nil
}
