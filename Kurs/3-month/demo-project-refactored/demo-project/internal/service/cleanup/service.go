package cleanup

import (
	"context"
	"errors"
	"fmt"
	"time"

	"demo-project/internal/models"
	"demo-project/internal/repository"
)

// Service удаляет мусор, который копится со временем. Расписание задаёт internal/cron.
type Service struct {
	files         repository.FileRepository
	storage       repository.ObjectStorage
	sessions      repository.SessionRepository
	unattachedTTL time.Duration
}

func NewService(
	files repository.FileRepository,
	storage repository.ObjectStorage,
	sessions repository.SessionRepository,
	unattachedTTL time.Duration,
) *Service {
	return &Service{
		files:         files,
		storage:       storage,
		sessions:      sessions,
		unattachedTTL: unattachedTTL,
	}
}

// DeleteUnattachedFiles удаляет файлы, которые загрузили, но так и не прикрепили к посту.
func (s *Service) DeleteUnattachedFiles(ctx context.Context) (int, error) {
	files, err := s.files.ListUnattachedBefore(ctx, time.Now().Add(-s.unattachedTTL))
	if err != nil {
		return 0, fmt.Errorf("list unattached files: %w", err)
	}

	deleted := 0
	for _, file := range files {
		err := s.files.Delete(ctx, file.ID)
		if errors.Is(err, models.ErrNotFound) {
			continue // файл успели прикрепить к посту — он больше не мусор
		}
		if err != nil {
			return deleted, fmt.Errorf("delete file %d: %w", file.ID, err)
		}
		if err := s.storage.Delete(ctx, file.ObjectKey); err != nil {
			return deleted, fmt.Errorf("delete file object %s: %w", file.ObjectKey, err)
		}
		deleted++
	}
	return deleted, nil
}

// DeleteExpiredSessions удаляет сессии, срок которых прошёл.
// Их токены в Redis к этому времени уже удалены по TTL.
func (s *Service) DeleteExpiredSessions(ctx context.Context) (int64, error) {
	deleted, err := s.sessions.DeleteExpired(ctx)
	if err != nil {
		return 0, fmt.Errorf("delete expired sessions: %w", err)
	}
	return deleted, nil
}
