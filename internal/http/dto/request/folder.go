package request

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type CreateFolderRequest struct {
	ConnectorID   string      `json:"connector_id" validate:"required"`
	DirectoryPath string      `json:"directory_path" validate:"required"`
	Tags          TagsRequest `json:"tags"`
}

type UpdateFolderRequest struct {
	ConnectorID   string      `json:"connector_id"`
	DirectoryPath string      `json:"directory_path" validate:"required"`
	Tags          TagsRequest `json:"tags"`
}

func (r *CreateFolderRequest) ToDomain() (*model.Folder, error) {
	folder := &model.Folder{
		Base:          model.NewBase(),
		ConnectorID:   r.ConnectorID,
		DirectoryPath: r.DirectoryPath,
	}

	// Convert TagsRequest to model.Tags
	tags := make(model.Tags, len(r.Tags))
	for i, tag := range r.Tags {
		tags[i] = model.Tag{
			Name:  tag.Name,
			Value: tag.Value,
		}
	}
	folder.Tags = tags

	if err := folder.Validate(); err != nil {
		return nil, err
	}

	return folder, nil
}

func (r *UpdateFolderRequest) ToDomain(id string) (*model.Folder, error) {
	folder, err := (&CreateFolderRequest{
		ConnectorID:   r.ConnectorID,
		DirectoryPath: r.DirectoryPath,
		Tags:          r.Tags,
	}).ToDomain()
	if err != nil {
		return nil, err
	}
	folder.ID = id
	return folder, nil
}
