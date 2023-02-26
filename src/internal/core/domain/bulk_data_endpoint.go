package domain

import "time"

type BulkDataEndpoint struct {
	ID          string    `json:"id"`
	Kind        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
	Size        int64     `json:"size"`
	DownloadURL string    `json:"download_uri"`
}
