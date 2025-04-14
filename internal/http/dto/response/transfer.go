package response

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type TransferResponse struct {
	ID                     string       `json:"id"`
	SourceFolderID         string       `json:"source_folder_id"`
	DestinationFolderID    string       `json:"destination_folder_id"`
	PostTransferSourcePath string       `json:"post_transfer_source_path"`
	Tags                   TagsResponse `json:"tags"`
	CreatedAt              string       `json:"created_at"`
	UpdatedAt              string       `json:"updated_at"`
}

func NewTransferResponse(transfer *model.Transfer) *TransferResponse {
	// Convert model.Tags to TagsResponse
	tags := make(TagsResponse, len(transfer.Tags))
	for i, tag := range transfer.Tags {
		tags[i] = TagResponse{
			Name:  tag.Name,
			Value: tag.Value,
		}
	}

	return &TransferResponse{
		ID:                     transfer.ID,
		SourceFolderID:         transfer.SourceFolderID,
		DestinationFolderID:    transfer.DestinationFolderID,
		PostTransferSourcePath: transfer.PostTransferSourcePath,
		Tags:                   tags,
		CreatedAt:              transfer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:              transfer.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
