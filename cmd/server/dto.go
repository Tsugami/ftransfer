package main

import "github.com/Tsugami/ftransfer/internal/storage_provider"

type CreateStorageProviderRequest struct {
	Name               string                 `json:"name" binding:"required" minlength:"1"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" binding:"required"`
	FileSystem         string                 `json:"file_system" binding:"required,oneof=UNIX WINDOWS"`
}

type UpdateStorageProviderRequest struct {
	Name               string                 `json:"name" binding:"required" minlength:"1"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" binding:"required"`
	FileSystem         string                 `json:"file_system" binding:"required,oneof=UNIX WINDOWS"`
}

type StorageProviderResponse struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	FileSystem         string                 `json:"file_system"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection"`
}

func NewStorageProviderResponse(storage *storage_provider.StorageProvider) StorageProviderResponse {
	return StorageProviderResponse{
		ID:                 storage.ID.String(),
		Name:               storage.Name,
		FileSystem:         string(storage.FileSystem),
		ProtocolConnection: storage.ProtocolConnection.GetJson(),
	}
}

func NewStorageProviderResponseList(storage []*storage_provider.StorageProvider) []StorageProviderResponse {
	response := make([]StorageProviderResponse, len(storage))
	for i, s := range storage {
		response[i] = NewStorageProviderResponse(s)
	}
	return response
}
