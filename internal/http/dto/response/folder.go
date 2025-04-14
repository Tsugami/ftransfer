package response

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type FolderResponse struct {
	ID            string       `json:"id"`
	ConnectorID   string       `json:"connector_id"`
	DirectoryPath string       `json:"directory_path"`
	Tags          TagsResponse `json:"tags"`
	CreatedAt     string       `json:"created_at"`
	UpdatedAt     string       `json:"updated_at"`
}

func NewFolderResponse(folder *model.Folder) *FolderResponse {
	// Convert model.Tags to TagsResponse
	tags := make(TagsResponse, len(folder.Tags))
	for i, tag := range folder.Tags {
		tags[i] = TagResponse{
			Name:  tag.Name,
			Value: tag.Value,
		}
	}

	return &FolderResponse{
		ID:            folder.ID,
		ConnectorID:   folder.ConnectorID,
		DirectoryPath: folder.DirectoryPath,
		Tags:          tags,
		CreatedAt:     folder.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:     folder.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
