package session

import (
	"context"
	"fmt"
	"net/netip"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"demo-project/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, s *models.Session) error {
	err := r.db.QueryRow(ctx,
		`INSERT INTO sessions (id, user_id, device_type, ip, user_agent, expires_at)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING created_at`,
		s.ID, s.UserID, s.DeviceType, ipOrNil(s.IP), s.UserAgent, s.ExpiresAt,
	).Scan(&s.CreatedAt)
	if err != nil {
		return fmt.Errorf("insert session: %w", err)
	}
	return nil
}

func (r *Repository) List(ctx context.Context, userID int64) ([]models.Session, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, device_type, ip, user_agent, created_at, expires_at
		   FROM sessions
		  WHERE user_id = $1 AND expires_at > now()
		  ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("select sessions: %w", err)
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan sessions: %w", err)
	}

	sessions := make([]models.Session, 0, len(items))
	for _, item := range items {
		sessions = append(sessions, item.toModel())
	}
	return sessions, nil
}

func (r *Repository) Extend(ctx context.Context, id string, expiresAt time.Time) error {
	tag, err := r.db.Exec(ctx, `UPDATE sessions SET expires_at = $2 WHERE id = $1`, id, expiresAt)
	if err != nil {
		return fmt.Errorf("extend session: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, userID int64, id string) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM sessions WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

// DeleteOthers удаляет все сессии пользователя, кроме keepID, и возвращает их id,
// чтобы сервис удалил и токены этих сессий.
func (r *Repository) DeleteOthers(ctx context.Context, userID int64, keepID string) ([]string, error) {
	rows, err := r.db.Query(ctx,
		`DELETE FROM sessions WHERE user_id = $1 AND id <> $2 RETURNING id`,
		userID, keepID,
	)
	if err != nil {
		return nil, fmt.Errorf("delete other sessions: %w", err)
	}

	ids, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("scan session ids: %w", err)
	}
	return ids, nil
}

func (r *Repository) DeleteExpired(ctx context.Context) (int64, error) {
	tag, err := r.db.Exec(ctx, `DELETE FROM sessions WHERE expires_at < now()`)
	if err != nil {
		return 0, fmt.Errorf("delete expired sessions: %w", err)
	}
	return tag.RowsAffected(), nil
}

// ipOrNil превращает пустой адрес в NULL: в колонке ip нельзя хранить «ничего» иначе.
func ipOrNil(ip netip.Addr) *netip.Addr {
	if !ip.IsValid() {
		return nil
	}
	return &ip
}
