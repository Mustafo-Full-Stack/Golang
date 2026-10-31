package file

import "demo-project/internal/models"

type dbModel struct {
	ID           int64  `db:"id"`
	UserID       int64  `db:"user_id"`
	PostID       *int64 `db:"post_id"`
	ObjectKey    string `db:"object_key"`
	OriginalName string `db:"original_name"`
	ContentType  string `db:"content_type"`
	SizeBytes    int64  `db:"size_bytes"`
}

func (m dbModel) toModel() models.File {
	return models.File{
		ID:          m.ID,
		UserID:      m.UserID,
		PostID:      m.PostID,
		ObjectKey:   m.ObjectKey,
		Name:        m.OriginalName,
		ContentType: m.ContentType,
		SizeBytes:   m.SizeBytes,
	}
}
