package habit

import (
	"context"
	"time"

	"habit-tracker/internal/models"
	habitrepo "habit-tracker/internal/repository/habit"
)

type Service struct {
	habits *habitrepo.Repository
}

func NewService(habits *habitrepo.Repository) *Service {
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

func (s *Service) List(ctx context.Context, userID int64) ([]models.Habit, error) {
	return s.habits.ListByUser(ctx, userID)
}

func (s *Service) Check(ctx context.Context, userID, habitID int64) error {
	if err := s.checkOwner(ctx, userID, habitID); err != nil {
		return err
	}

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return s.habits.CreateCheck(ctx, habitID, today)
}

func (s *Service) Checks(ctx context.Context, userID, habitID int64) ([]models.HabitCheck, error) {
	if err := s.checkOwner(ctx, userID, habitID); err != nil {
		return nil, err
	}
	return s.habits.ListChecks(ctx, habitID)
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
