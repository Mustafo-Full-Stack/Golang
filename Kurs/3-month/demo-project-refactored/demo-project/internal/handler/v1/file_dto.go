package v1

import "demo-project/internal/models"

type fileResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	SizeBytes   int64  `json:"size_bytes"`
	URL         string `json:"url"`
}

func newFileResponse(f *models.File) fileResponse {
	return fileResponse{
		ID:          f.ID,
		Name:        f.Name,
		ContentType: f.ContentType,
		SizeBytes:   f.SizeBytes,
		URL:         f.URL,
	}
}

func newFilesResponse(files []models.File) []fileResponse {
	resp := make([]fileResponse, 0, len(files))
	for i := range files {
		resp = append(resp, newFileResponse(&files[i]))
	}
	return resp
}
