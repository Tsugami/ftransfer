package request

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type CreateTransferRequest struct {
	SourceFolderID         string      `json:"source_folder_id" validate:"required"`
	DestinationFolderID    string      `json:"destination_folder_id" validate:"required"`
	PostTransferSourcePath string      `json:"post_transfer_source_path" validate:"required"`
	Tags                   TagsRequest `json:"tags"`
}

type UpdateTransferRequest struct {
	SourceFolderID         string      `json:"source_folder_id"`
	DestinationFolderID    string      `json:"destination_folder_id"`
	PostTransferSourcePath string      `json:"post_transfer_source_path"`
	Tags                   TagsRequest `json:"tags"`
}

func (r *CreateTransferRequest) ToDomain() (*model.Transfer, error) {
	base := model.NewBase()

	// Convert TagsRequest to model.Tags
	tags := make(model.Tags, len(r.Tags))
	for i, tag := range r.Tags {
		tags[i] = model.Tag{
			Name:  tag.Name,
			Value: tag.Value,
		}
	}
	base.Tags = tags

	transfer := &model.Transfer{
		Base:                   base,
		SourceFolderID:         r.SourceFolderID,
		DestinationFolderID:    r.DestinationFolderID,
		PostTransferSourcePath: r.PostTransferSourcePath,
	}

	if err := transfer.Validate(); err != nil {
		return nil, err
	}

	return transfer, nil
}

func (r *UpdateTransferRequest) ToDomain(id string) (*model.Transfer, error) {
	transfer, err := (&CreateTransferRequest{
		SourceFolderID:         r.SourceFolderID,
		DestinationFolderID:    r.DestinationFolderID,
		PostTransferSourcePath: r.PostTransferSourcePath,
		Tags:                   r.Tags,
	}).ToDomain()
	if err != nil {
		return nil, err
	}
	transfer.ID = id
	return transfer, nil
}
