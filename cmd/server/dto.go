package main

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
